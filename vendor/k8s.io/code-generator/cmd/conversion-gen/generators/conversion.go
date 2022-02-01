/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generators

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"

	"k8s.io/klog/v2"

	conversionargs "k8s.io/code-generator/cmd/conversion-gen/args"
	genutil "k8s.io/code-generator/pkg/util"
)

// These are the comment tags that carry parameters for conversion generation.
const (
	// In doc.go:
	//
	//   "+k8s:conversion-gen=<peer-pkg>" indicates that  <peer-pkg> is the
	//   import path of the package the peer types are defined in.
	//
	// On a type:
	//
	//   "+k8s:conversion-gen=false" on type <src> causes conversion-gen to
	//   not attempt to generate any "Convert_<src>_To_<dst>" functions.  It
	//   may still generate other functions that call such a function if
	//   necessary; in that case it would be the user's responsibility to
	//   manually write the function.
	//
	// On a struct member:
	//
	//   "+k8s:conversion-gen=false" on a member within <src> struct causes
	//   the generated "Convert_<src>_To_<dst>" functions to ignore that
	//   field, not attempting to set <dst>.<fieldname>.
	tagName = "k8s:conversion-gen"
	// On a type:
	//
	//   "+k8s:conversion-gen:explicit-from=net/url.Values" results in
	//   generating conversion from net/url.Values.  This is a special case;
	//   other types are not supported.
	explicitFromTagName = "k8s:conversion-gen:explicit-from"
	// In doc.go:
	//
	//   "+k8s:conversion-gen-external-types=<type-pkg>" indicates that
	//   <type-pkg> is the relative path to the package the types are
	//   defined in.
	externalTypesTagName = "k8s:conversion-gen-external-types"
	// On a "Convert_<src>_To_<dst>" function:
	//
	//   "+k8s:conversion-fn=copy-only" indicates that the function is
	//   merely blindly copying data (not doing any higher-level conversion
	//   logic), and that it is OK to optimize the function away if
	//   conversion-gen determines the types to be memory-equivalent.
	//
	//   "+k8s:conversion-fn=drop" indicates that the conversion done by
	//   this function should never be performed on struct members; as if
	//   "+k8s:conversion-gen=false" were added to the source struct member.
	conversionFnTagName = "k8s:conversion-fn"
	// On a struct member:
	//
	//   "+k8s:conversion-gen:rename=<peer-name>" indicates that the
	//   equivalent field in the equivalent type in the <peer-pkg> is named
	//   <peer-name>.  This applies in both directions; e.g. setting this
	//   marker on a field in pkg1.Foo will affect both
	//   "Convert_pkg1_Foo_To_pkg2_Foo" and "Convert_pkg2_Foo_To_pkg1_Foo";
	//   it is not necessary to set in both packages.  This will not prevent
	//   a field from pairing with an identically named field; it does not
	//   mask the original name, it adds to it.  As such, it may be given
	//   multiple times, in order to facilitate generating conversions with
	//   many different peer packages.
	renameTagName = "k8s:conversion-gen:rename"
)

func extractTag(comments []string) []string {
	return types.ExtractCommentTags("+", comments)[tagName]
}

func extractExplicitFromTag(comments []string) []string {
	return types.ExtractCommentTags("+", comments)[explicitFromTagName]
}

func extractExternalTypesTag(comments []string) []string {
	return types.ExtractCommentTags("+", comments)[externalTypesTagName]
}

func isCopyOnly(comments []string) bool {
	values := types.ExtractCommentTags("+", comments)[conversionFnTagName]
	return len(values) == 1 && values[0] == "copy-only"
}

func isDrop(comments []string) bool {
	values := types.ExtractCommentTags("+", comments)[conversionFnTagName]
	return len(values) == 1 && values[0] == "drop"
}

func extractRenameTags(comments []string) []string {
	return types.ExtractCommentTags("+", comments)[renameTagName]
}

// TODO: This is created only to reduce number of changes in a single PR.
// Remove it and use PublicNamer instead.
func conversionNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Join: func(pre string, in []string, post string) string {
			return strings.Join(in, "_")
		},
		PrependPackageNames: 1,
	}
}

func defaultFnNamer() *namer.NameStrategy {
	return &namer.NameStrategy{
		Prefix: "SetDefaults_",
		Join: func(pre string, in []string, post string) string {
			return pre + strings.Join(in, "_") + post
		},
	}
}

// NameSystems returns the name system used by the generators in this package.
func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public":    conversionNamer(),
		"raw":       namer.NewRawNamer("", nil),
		"defaultfn": defaultFnNamer(),
	}
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func DefaultNameSystem() string {
	return "public"
}

func getPeerTypeFor(context *generator.Context, t *types.Type, potenialPeerPkgs []string) *types.Type {
	for _, ppp := range potenialPeerPkgs {
		p := context.Universe.Package(ppp)
		if p == nil {
			continue
		}
		if p.Has(t.Name.Name) {
			return p.Type(t.Name.Name)
		}
	}
	return nil
}

type conversionPair struct {
	inType  *types.Type
	outType *types.Type
}

// All of the types in conversions map are of type "DeclarationOf" with
// the underlying type being "Func".
type conversionFuncMap map[conversionPair]*types.Type

// Returns all manually-defined conversion functions in the package.
func getManualConversionFunctions(context *generator.Context, pkg *types.Package, manualMap conversionFuncMap) {
	if pkg == nil {
		klog.Warningf("Skipping nil package passed to getManualConversionFunctions")
		return
	}
	klog.V(5).Infof("Scanning for conversion functions in %v", pkg.Name)

	scopeName := types.Ref(conversionPackagePath, "Scope").Name
	errorName := types.Ref("", "error").Name
	buffer := &bytes.Buffer{}
	sw := generator.NewSnippetWriter(buffer, context, "$", "$")

	for _, f := range pkg.Functions {
		if f.Underlying == nil || f.Underlying.Kind != types.Func {
			klog.Errorf("Malformed function: %#v", f)
			continue
		}
		if f.Underlying.Signature == nil {
			klog.Errorf("Function without signature: %#v", f)
			continue
		}
		klog.V(8).Infof("Considering function %s", f.Name)
		signature := f.Underlying.Signature
		// Check whether the function is conversion function.
		// Note that all of them have signature:
		// func Convert_inType_To_outType(inType, outType, conversion.Scope) error
		if signature.Receiver != nil {
			klog.V(8).Infof("%s has a receiver", f.Name)
			continue
		}
		if len(signature.Parameters) != 3 || signature.Parameters[2].Name != scopeName {
			klog.V(8).Infof("%s has wrong parameters", f.Name)
			continue
		}
		if len(signature.Results) != 1 || signature.Results[0].Name != errorName {
			klog.V(8).Infof("%s has wrong results", f.Name)
			continue
		}
		inType := signature.Parameters[0]
		outType := signature.Parameters[1]
		if inType.Kind != types.Pointer || outType.Kind != types.Pointer {
			klog.V(8).Infof("%s has wrong parameter types", f.Name)
			continue
		}
		// Now check if the name satisfies the convention.
		// TODO: This should call the Namer directly.
		args := argsFromType(inType.Elem, outType.Elem)
		sw.Do("Convert_$.inType|public$_To_$.outType|public$", args)
		if f.Name.Name == buffer.String() {
			klog.V(4).Infof("Found conversion function %s", f.Name)
			key := conversionPair{inType.Elem, outType.Elem}
			// We might scan the same package twice, and that's OK.
			if v, ok := manualMap[key]; ok && v != nil && v.Name.Package != pkg.Path {
				panic(fmt.Sprintf("duplicate static conversion defined: %s -> %s from:\n%s.%s\n%s.%s", key.inType, key.outType, v.Name.Package, v.Name.Name, f.Name.Package, f.Name.Name))
			}
			manualMap[key] = f
		} else {
			// prevent user error when they don't get the correct conversion signature
			if strings.HasPrefix(f.Name.Name, "Convert_") {
				klog.Errorf("Rename function %s %s -> %s to match expected conversion signature", f.Name.Package, f.Name.Name, buffer.String())
			}
			klog.V(8).Infof("%s has wrong name", f.Name)
		}
		buffer.Reset()
	}
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	boilerplate, err := arguments.LoadGoBoilerplate()
	if err != nil {
		klog.Fatalf("Failed loading boilerplate: %v", err)
	}

	packages := generator.Packages{}
	header := append([]byte(fmt.Sprintf("// +build !%s\n\n", arguments.GeneratedBuildTag)), boilerplate...)

	// Accumulate pre-existing conversion functions.
	// TODO: This is too ad-hoc.  We need a better way.
	manualConversions := conversionFuncMap{}

	// Record types that are memory equivalent. A type is memory equivalent
	// if it has the same memory layout and no nested manual conversion is
	// defined.
	// TODO: in the future, relax the nested manual conversion requirement
	//   if we can show that a large enough types are memory identical but
	//   have non-trivial conversion
	memoryEquivalentTypes := equalMemoryTypes{}

	// We are generating conversions only for packages that are explicitly
	// passed as InputDir.
	processed := map[string]bool{}
	for _, i := range context.Inputs {
		// skip duplicates
		if processed[i] {
			continue
		}
		processed[i] = true

		klog.V(5).Infof("considering pkg %q", i)
		pkg := context.Universe[i]
		// typesPkg is where the versioned types are defined. Sometimes it is
		// different from pkg. For example, kubernetes core/v1 types are defined
		// in vendor/k8s.io/api/core/v1, while pkg is at pkg/api/v1.
		typesPkg := pkg
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}

		// Add conversion and defaulting functions.
		getManualConversionFunctions(context, pkg, manualConversions)

		// Only generate conversions for packages which explicitly request it
		// by specifying one or more "+k8s:conversion-gen=<peer-pkg>"
		// in their doc.go file.
		peerPkgs := extractTag(pkg.Comments)
		if peerPkgs != nil {
			klog.V(5).Infof("  tags: %q", peerPkgs)
			if len(peerPkgs) == 1 && peerPkgs[0] == "false" {
				// If a single +k8s:conversion-gen=false tag is defined, we still want
				// the generator to fire for this package for explicit conversions, but
				// we are clearing the peerPkgs to not generate any standard conversions.
				peerPkgs = nil
			}
		} else {
			klog.V(5).Infof("  no tag")
			continue
		}
		skipUnsafe := false
		extraDirs := []string{}
		if customArgs, ok := arguments.CustomArgs.(*conversionargs.CustomArgs); ok {
			if len(peerPkgs) > 0 {
				peerPkgs = append(peerPkgs, customArgs.BasePeerDirs...)
				peerPkgs = append(peerPkgs, customArgs.ExtraPeerDirs...)
			}
			extraDirs = customArgs.ExtraDirs
			skipUnsafe = customArgs.SkipUnsafe
		}

		// if the external types are not in the same package where the conversion functions to be generated
		externalTypesValues := extractExternalTypesTag(pkg.Comments)
		if externalTypesValues != nil {
			if len(externalTypesValues) != 1 {
				klog.Fatalf("  expect only one value for %q tag, got: %q", externalTypesTagName, externalTypesValues)
			}
			externalTypes := externalTypesValues[0]
			klog.V(5).Infof("  external types tags: %q", externalTypes)
			var err error
			typesPkg, err = context.AddDirectory(externalTypes)
			if err != nil {
				klog.Fatalf("cannot import package %s", externalTypes)
			}
			// update context.Order to the latest context.Universe
			orderer := namer.Orderer{Namer: namer.NewPublicNamer(1)}
			context.Order = orderer.OrderUniverse(context.Universe)
		}

		// if the source path is within a /vendor/ directory (for example,
		// k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1), allow
		// generation to output to the proper relative path (under vendor).
		// Otherwise, the generator will create the file in the wrong location
		// in the output directory.
		// TODO: build a more fundamental concept in gengo for dealing with modifications
		// to vendored packages.
		for i := range peerPkgs {
			peerPkgs[i] = genutil.Vendorless(peerPkgs[i])
		}
		for i := range extraDirs {
			extraDirs[i] = genutil.Vendorless(extraDirs[i])
		}

		// Make sure our peer-packages are added and fully parsed.
		for _, pp := range append(peerPkgs, extraDirs...) {
			context.AddDir(pp)
			p := context.Universe[pp]
			if nil == p {
				klog.Fatalf("failed to find pkg: %s", pp)
			}
			getManualConversionFunctions(context, p, manualConversions)
		}

		unsafeEquality := TypesEqual(memoryEquivalentTypes)
		if skipUnsafe {
			unsafeEquality = noEquality{}
		}

		path := pkg.Path
		// if the source path is within a /vendor/ directory (for example,
		// k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1), allow
		// generation to output to the proper relative path (under vendor).
		// Otherwise, the generator will create the file in the wrong location
		// in the output directory.
		// TODO: build a more fundamental concept in gengo for dealing with modifications
		// to vendored packages.
		if strings.HasPrefix(pkg.SourcePath, arguments.OutputBase) {
			expandedPath := strings.TrimPrefix(pkg.SourcePath, arguments.OutputBase)
			if strings.Contains(expandedPath, "/vendor/") {
				path = expandedPath
			}
		}
		packages = append(packages,
			&generator.DefaultPackage{
				PackageName: filepath.Base(pkg.Path),
				PackagePath: path,
				HeaderText:  header,
				GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
					return []generator.Generator{
						NewGenConversion(arguments.OutputFileBaseName, typesPkg.Path, pkg.Path, manualConversions, peerPkgs, unsafeEquality),
					}
				},
				FilterFunc: func(c *generator.Context, t *types.Type) bool {
					return t.Name.Package == typesPkg.Path
				},
			})
	}

	return packages
}

type equalMemoryTypes map[conversionPair]bool

func (e equalMemoryTypes) Equal(a, b *types.Type) bool {
	// alreadyVisitedTypes holds all the types that have already been checked in the structural type recursion.
	alreadyVisitedTypes := make(map[*types.Type]bool)
	return e.cachingEqual(a, b, alreadyVisitedTypes)
}

func (e equalMemoryTypes) cachingEqual(a, b *types.Type, alreadyVisitedTypes map[*types.Type]bool) bool {
	if a == b {
		return true
	}
	if equal, ok := e[conversionPair{a, b}]; ok {
		return equal
	}
	if equal, ok := e[conversionPair{b, a}]; ok {
		return equal
	}
	result := e.equal(a, b, alreadyVisitedTypes)
	e[conversionPair{a, b}] = result
	e[conversionPair{b, a}] = result
	return result
}

func (e equalMemoryTypes) equal(a, b *types.Type, alreadyVisitedTypes map[*types.Type]bool) bool {
	in, out := unwrapAlias(a), unwrapAlias(b)
	switch {
	case in == out:
		return true
	case in.Kind == out.Kind:
		// if the type exists already, return early to avoid recursion
		if alreadyVisitedTypes[in] {
			return true
		}
		alreadyVisitedTypes[in] = true

		switch in.Kind {
		case types.Struct:
			if len(in.Members) != len(out.Members) {
				return false
			}
			for i, inMember := range in.Members {
				outMember := out.Members[i]
				if !e.cachingEqual(inMember.Type, outMember.Type, alreadyVisitedTypes) {
					return false
				}
			}
			return true
		case types.Pointer:
			return e.cachingEqual(in.Elem, out.Elem, alreadyVisitedTypes)
		case types.Map:
			return e.cachingEqual(in.Key, out.Key, alreadyVisitedTypes) && e.cachingEqual(in.Elem, out.Elem, alreadyVisitedTypes)
		case types.Slice:
			return e.cachingEqual(in.Elem, out.Elem, alreadyVisitedTypes)
		case types.Interface:
			// TODO: determine whether the interfaces are actually equivalent - for now, they must have the
			// same type.
			return false
		case types.Builtin:
			return in.Name.Name == out.Name.Name
		}
	}
	return false
}

func findMember(t *types.Type, names ...string) (types.Member, bool) {
	if t.Kind != types.Struct {
		return types.Member{}, false
	}
	for _, member := range t.Members {
		for _, name := range names {
			if member.Name == name {
				return member, true
			}
			for _, memberName := range extractRenameTags(member.CommentLines) {
				if memberName == name {
					return member, true
				}
			}
		}
	}
	return types.Member{}, false
}

// unwrapAlias recurses down aliased types to find the bedrock type.
func unwrapAlias(in *types.Type) *types.Type {
	for in.Kind == types.Alias {
		in = in.Underlying
	}
	return in
}

const (
	runtimePackagePath    = "k8s.io/apimachinery/pkg/runtime"
	conversionPackagePath = "k8s.io/apimachinery/pkg/conversion"
)

type noEquality struct{}

func (noEquality) Equal(_, _ *types.Type) bool { return false }

type TypesEqual interface {
	Equal(a, b *types.Type) bool
}

// genConversion produces a file with a autogenerated conversions.
type genConversion struct {
	generator.DefaultGen
	// the package that contains the types that conversion func are going to be
	// generated for
	typesPackage string
	// the package that the conversion funcs are going to be output to
	outputPackage string
	// packages that contain the peer of types in typesPacakge
	peerPackages        []string
	manualConversions   conversionFuncMap
	imports             namer.ImportTracker
	types               []*types.Type
	explicitConversions []conversionPair
	skippedFields       map[*types.Type][]string
	useUnsafe           TypesEqual
}

func NewGenConversion(sanitizedName, typesPackage, outputPackage string, manualConversions conversionFuncMap, peerPkgs []string, useUnsafe TypesEqual) generator.Generator {
	return &genConversion{
		DefaultGen: generator.DefaultGen{
			OptionalName: sanitizedName,
		},
		typesPackage:        typesPackage,
		outputPackage:       outputPackage,
		peerPackages:        peerPkgs,
		manualConversions:   manualConversions,
		imports:             generator.NewImportTracker(),
		types:               []*types.Type{},
		explicitConversions: []conversionPair{},
		skippedFields:       map[*types.Type][]string{},
		useUnsafe:           useUnsafe,
	}
}

func (g *genConversion) Namers(c *generator.Context) namer.NameSystems {
	// Have the raw namer for this file track what it imports.
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
		"publicIT": &namerPlusImportTracking{
			delegate: conversionNamer(),
			tracker:  g.imports,
		},
	}
}

type namerPlusImportTracking struct {
	delegate namer.Namer
	tracker  namer.ImportTracker
}

func (n *namerPlusImportTracking) Name(t *types.Type) string {
	n.tracker.AddType(t)
	return n.delegate.Name(t)
}

// convertibleOnlyWithinPackage returns whether the conversion between
// inType and outType is something that we "want" a
// "Convert_{inType}_To_{outType}" function for, whether that function
// is able to be generated by conversion-gen or needs to be manually
// provided.
//
// These criteria are:
//  1. One of the types is in the package that we're generating conversions for, and
//  2. that type has not opted out of conversion with +k8s:conversion-gen=false, and
//  3. both types are exported named types, and
//  4. both types are (after resolving aliases) of the same kind, and
//  5. that kind is a Builtin, Map, Slice, Struct, or Pointer.
func (g *genConversion) convertibleOnlyWithinPackage(inType, outType *types.Type) bool {
	var t *types.Type
	var other *types.Type
	if inType.Name.Package == g.typesPackage {
		t, other = inType, outType
	} else {
		t, other = outType, inType
	}

	// 1.
	if t.Name.Package != g.typesPackage {
		return false
	}

	// 2. If the type has opted out, skip it.
	tagvals := extractTag(t.CommentLines)
	if tagvals != nil {
		if tagvals[0] != "false" {
			klog.Fatalf("Type %v: unsupported %s value: %q", t, tagName, tagvals[0])
		}
		klog.V(5).Infof("type %v requests no conversion generation, skipping", t)
		return false
	}

	// 3. Filter out private types.
	if namer.IsPrivateGoName(t.Name.Name) || namer.IsPrivateGoName(other.Name.Name) {
		return false
	}

	// 4. If the types aren't of the same kind, skip it.
	t, other = unwrapAlias(t), unwrapAlias(other)
	if t.Kind != other.Kind {
		return false
	}

	// 5. If that kind isn't one that we support, skip it.
	switch t.Kind {
	case types.Builtin, types.Map, types.Slice, types.Struct, types.Pointer:
		// ok
	default:
		// skip it
		return false
	}

	return true
}

func getExplicitFromTypes(t *types.Type) []types.Name {
	comments := append(t.SecondClosestCommentLines, t.CommentLines...)
	paths := extractExplicitFromTag(comments)
	result := []types.Name{}
	for _, path := range paths {
		items := strings.Split(path, ".")
		if len(items) != 2 {
			klog.Errorf("Unexpected k8s:conversion-gen:explicit-from tag: %s", path)
			continue
		}
		switch {
		case items[0] == "net/url" && items[1] == "Values":
		default:
			klog.Fatalf("Not supported k8s:conversion-gen:explicit-from tag: %s", path)
		}
		result = append(result, types.Name{Package: items[0], Name: items[1]})
	}
	return result
}

func (g *genConversion) Filter(c *generator.Context, t *types.Type) bool {
	convertibleWithPeer := func() bool {
		peerType := getPeerTypeFor(c, t, g.peerPackages)
		if peerType == nil {
			return false
		}
		if !g.convertibleOnlyWithinPackage(t, peerType) {
			return false
		}
		g.types = append(g.types, t)
		return true
	}()

	explicitlyConvertible := func() bool {
		inTypes := getExplicitFromTypes(t)
		if len(inTypes) == 0 {
			return false
		}
		for i := range inTypes {
			pair := conversionPair{
				inType:  &types.Type{Name: inTypes[i]},
				outType: t,
			}
			g.explicitConversions = append(g.explicitConversions, pair)
		}
		return true
	}()

	return convertibleWithPeer || explicitlyConvertible
}

func (g *genConversion) isOtherPackage(pkg string) bool {
	if pkg == g.outputPackage {
		return false
	}
	if strings.HasSuffix(pkg, `"`+g.outputPackage+`"`) {
		return false
	}
	return true
}

func (g *genConversion) Imports(c *generator.Context) (imports []string) {
	var importLines []string
	for _, singleImport := range g.imports.ImportLines() {
		if g.isOtherPackage(singleImport) {
			importLines = append(importLines, singleImport)
		}
	}
	return importLines
}

func argsFromType(inType, outType *types.Type) generator.Args {
	return generator.Args{
		"inType":  inType,
		"outType": outType,
	}
}

const nameTmpl = "Convert_$.inType|publicIT$_To_$.outType|publicIT$"

func (g *genConversion) preexists(inType, outType *types.Type) (*types.Type, bool) {
	function, ok := g.manualConversions[conversionPair{inType, outType}]
	return function, ok
}

func (g *genConversion) Init(c *generator.Context, w io.Writer) error {
	if klog.V(5).Enabled() {
		if m, ok := g.useUnsafe.(equalMemoryTypes); ok {
			var result []string
			klog.Infof("All objects without identical memory layout:")
			for k, v := range m {
				if v {
					continue
				}
				result = append(result, fmt.Sprintf("  %s -> %s = %t", k.inType, k.outType, v))
			}
			sort.Strings(result)
			for _, s := range result {
				klog.Infof(s)
			}
		}
	}
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	sw.Do("func init() {\n", nil)
	sw.Do("localSchemeBuilder.Register(RegisterConversions)\n", nil)
	sw.Do("}\n", nil)

	scheme := c.Universe.Type(types.Name{Package: runtimePackagePath, Name: "Scheme"})
	schemePtr := &types.Type{
		Kind: types.Pointer,
		Elem: scheme,
	}
	sw.Do("// RegisterConversions adds conversion functions to the given scheme.\n", nil)
	sw.Do("// Public to allow building arbitrary schemes.\n", nil)
	sw.Do("func RegisterConversions(s $.|raw$) error {\n", schemePtr)
	for _, t := range g.types {
		peerType := getPeerTypeFor(c, t, g.peerPackages)
		if _, found := g.preexists(t, peerType); !found {
			args := argsFromType(t, peerType).With("Scope", types.Ref(conversionPackagePath, "Scope"))
			sw.Do("if err := s.AddGeneratedConversionFunc((*$.inType|raw$)(nil), (*$.outType|raw$)(nil), func(a, b interface{}, scope $.Scope|raw$) error { return "+nameTmpl+"(a.(*$.inType|raw$), b.(*$.outType|raw$), scope) }); err != nil { return err }\n", args)
		}
		if _, found := g.preexists(peerType, t); !found {
			args := argsFromType(peerType, t).With("Scope", types.Ref(conversionPackagePath, "Scope"))
			sw.Do("if err := s.AddGeneratedConversionFunc((*$.inType|raw$)(nil), (*$.outType|raw$)(nil), func(a, b interface{}, scope $.Scope|raw$) error { return "+nameTmpl+"(a.(*$.inType|raw$), b.(*$.outType|raw$), scope) }); err != nil { return err }\n", args)
		}
	}

	for i := range g.explicitConversions {
		args := argsFromType(g.explicitConversions[i].inType, g.explicitConversions[i].outType).With("Scope", types.Ref(conversionPackagePath, "Scope"))
		sw.Do("if err := s.AddGeneratedConversionFunc((*$.inType|raw$)(nil), (*$.outType|raw$)(nil), func(a, b interface{}, scope $.Scope|raw$) error { return "+nameTmpl+"(a.(*$.inType|raw$), b.(*$.outType|raw$), scope) }); err != nil { return err }\n", args)
	}

	var pairs []conversionPair
	for pair, t := range g.manualConversions {
		if t.Name.Package != g.outputPackage {
			continue
		}
		pairs = append(pairs, pair)
	}
	// sort by name of the conversion function
	sort.Slice(pairs, func(i, j int) bool {
		if g.manualConversions[pairs[i]].Name.Name < g.manualConversions[pairs[j]].Name.Name {
			return true
		}
		return false
	})
	for _, pair := range pairs {
		args := argsFromType(pair.inType, pair.outType).With("Scope", types.Ref(conversionPackagePath, "Scope")).With("fn", g.manualConversions[pair])
		sw.Do("if err := s.AddConversionFunc((*$.inType|raw$)(nil), (*$.outType|raw$)(nil), func(a, b interface{}, scope $.Scope|raw$) error { return $.fn|raw$(a.(*$.inType|raw$), b.(*$.outType|raw$), scope) }); err != nil { return err }\n", args)
	}

	sw.Do("return nil\n", nil)
	sw.Do("}\n\n", nil)
	return sw.Error()
}

func (g *genConversion) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	klog.V(5).Infof("generating for type %v", t)
	sw := generator.NewSnippetWriter(w, c, "$", "$")

	if peerType := getPeerTypeFor(c, t, g.peerPackages); peerType != nil {
		g.generateConversion(t, peerType, sw)
		g.generateConversion(peerType, t, sw)
	}

	for _, inTypeName := range getExplicitFromTypes(t) {
		inPkg, ok := c.Universe[inTypeName.Package]
		if !ok {
			klog.Errorf("Unrecognized package: %s", inTypeName.Package)
			continue
		}
		inType, ok := inPkg.Types[inTypeName.Name]
		if !ok {
			klog.Errorf("Unrecognized type in package %s: %s", inTypeName.Package, inTypeName.Name)
			continue
		}
		switch {
		case inType.Name.Package == "net/url" && inType.Name.Name == "Values":
			g.generateFromUrlValues(inType, t, sw)
		default:
			klog.Errorf("Not supported input type: %#v", inType.Name)
		}
	}

	return sw.Error()
}

func (g *genConversion) generateConversion(inType, outType *types.Type, sw *generator.SnippetWriter) {
	args := argsFromType(inType, outType).
		With("Scope", types.Ref(conversionPackagePath, "Scope"))

	sw.Do("func auto"+nameTmpl+"(in *$.inType|raw$, out *$.outType|raw$, s $.Scope|raw$) error {\n", args)
	g.generateFor(inType, outType, false, sw)
	sw.Do("return nil\n", nil)
	sw.Do("}\n\n", nil)

	if _, found := g.preexists(inType, outType); found {
		// There is a public manual Conversion method: use it.
	} else if skipped := g.skippedFields[inType]; len(skipped) != 0 {
		// The inType had some fields we could not generate.
		klog.Errorf("Warning: could not find nor generate a final Conversion function for %v -> %v", inType, outType)
		klog.Errorf("  the following fields need manual conversion:")
		for _, f := range skipped {
			klog.Errorf("      - %v", f)
		}
	} else {
		// Emit a public conversion function.
		sw.Do("// "+nameTmpl+" is an autogenerated conversion function.\n", args)
		sw.Do("func "+nameTmpl+"(in *$.inType|raw$, out *$.outType|raw$, s $.Scope|raw$) error {\n", args)
		sw.Do("return auto"+nameTmpl+"(in, out, s)\n", args)
		sw.Do("}\n\n", nil)
	}
}

// we use the system of shadowing 'in' and 'out' so that the same code is valid
// at any nesting level. This makes the autogenerator easy to understand, and
// the compiler shouldn't care.
func (g *genConversion) generateFor(inType, outType *types.Type, functionOK bool, sw *generator.SnippetWriter) {
	klog.V(5).Infof("generating %v -> %v", inType, outType)

	inTypeResolved, outTypeResolved := inType, outType
	if underlying := unwrapAlias(inTypeResolved); underlying != inTypeResolved {
		copied := *underlying
		copied.Name = inTypeResolved.Name
		inTypeResolved = &copied
	}
	if underlying := unwrapAlias(outTypeResolved); underlying != outTypeResolved {
		copied := *underlying
		copied.Name = outTypeResolved.Name
		outTypeResolved = &copied
	}

	// (1) existing function
	if functionOK {
		if function, ok := g.preexists(inType, outType); ok {
			// As an example, conversion functions exist that allow types with private fields to be
			// correctly copied between types. These functions are equivalent to a memory assignment,
			// and are necessary for the reflection path, but should not block memory conversion.
			// Convert_unversioned_Time_to_unversioned_Time is an example of this logic.
			if isCopyOnly(function.CommentLines) && (isDirectlyConvertible(inType, outType) || g.useUnsafe.Equal(inTypeResolved, outTypeResolved)) {
				klog.V(5).Infof("Skipped function %s because it is copy-only and we can use direct assignment or unsafe casting", function.Name)
			} else {
				sw.Do("if err := $.|raw$(in, out, s); err != nil {\n", function)
				sw.Do("return err\n", nil)
				sw.Do("}\n", nil)
				return
			}
		} else if g.convertibleOnlyWithinPackage(inType, outType) {
			sw.Do("if err := "+nameTmpl+"(in, out, s); err != nil {\n", argsFromType(inType, outType))
			sw.Do("return err\n", nil)
			sw.Do("}\n", nil)
			return
		}
	}

	// (2) unsafe casting
	if g.useUnsafe.Equal(inTypeResolved, outTypeResolved) {
		args := argsFromType(inType, outType).
			With("Pointer", types.Ref("unsafe", "Pointer"))
		switch inTypeResolved.Kind {
		case types.Pointer:
			sw.Do("out = ($.outType|raw$)($.Pointer|raw$(in))\n", args)
			return
		case types.Map:
			sw.Do("out = *(*$.outType|raw$)($.Pointer|raw$(&in))\n", args)
			return
		case types.Slice:
			sw.Do("out = *(*$.outType|raw$)($.Pointer|raw$(&in))\n", args)
			return
		}
	}

	// (3) direct assignment
	if isDirectlyAssignable(inType, outType) {
		sw.Do("*out = *in\n", nil)
		return
	}

	// (4) direct conversion
	if isDirectlyConvertible(inType, outType) {
		sw.Do("*out = $.|raw$(*in)\n", outType)
		return
	}

	// (5) generate code
	if inTypeResolved.Kind == outTypeResolved.Kind {
		switch inTypeResolved.Kind {
		case types.Map:
			g.doMap(inTypeResolved, outTypeResolved, sw)
			return
		case types.Slice:
			g.doSlice(inTypeResolved, outTypeResolved, sw)
			return
		case types.Struct:
			g.doStruct(inTypeResolved, outTypeResolved, sw)
			return
		case types.Pointer:
			g.doPointer(inTypeResolved, outTypeResolved, sw)
			return
		}
	}

	// (6) fail
	g.doCompileErrorOnMissingConversion(inType, outType, sw)
}

func (g *genConversion) doCompileErrorOnMissingConversion(inType, outType *types.Type, sw *generator.SnippetWriter) {
	sw.Do("// FIXME: Provide conversion function to convert $.inType|raw$ to $.outType|raw$;\n",
		argsFromType(inType, outType))
	if len(g.manualConversions) == 0 {
		sw.Do("// no manual conversion functions are currently provided.\n", nil)
	} else {
		sw.Do("// the currently provided manual conversion functions are\n", nil)
		pairs := make([]conversionPair, 0, len(g.manualConversions))
		for pair := range g.manualConversions {
			pairs = append(pairs, pair)
		}
		sort.Slice(pairs, func(i, j int) bool {
			return g.manualConversions[pairs[i]].String() < g.manualConversions[pairs[j]].String()
		})
		for _, pair := range pairs {
			sw.Do("//  - $.func|raw$() ($.inType|raw$ to $.outType|raw$)\n", generator.Args{
				"inType":  pair.inType,
				"outType": pair.outType,
				"func":    g.manualConversions[pair],
			})
		}
	}
	sw.Do("compileErrorOnMissingConversion()\n", nil)
}

func (g *genConversion) doMap(inType, outType *types.Type, sw *generator.SnippetWriter) {
	sw.Do("if *in == nil {\n", nil)
	sw.Do("*out = nil\n", nil)
	sw.Do("} else {\n", nil)

	sw.Do("*out = make($.|raw$, len(*in))\n", outType)
	sw.Do("for inKey, inVal := range *in {\n", nil)

	sw.Do("outKey := new($.|raw$)\n", outType.Key)
	sw.Do("if true {\n", nil)
	sw.Do("in, out := &inKey, outKey\n", nil)
	g.generateFor(inType.Key, outType.Key, true, sw)
	sw.Do("}\n", nil)

	sw.Do("outVal := new($.|raw$)\n", outType.Elem)
	sw.Do("if true {\n", nil)
	sw.Do("in, out := &inVal, outVal\n", nil)
	g.generateFor(inType.Elem, outType.Elem, true, sw)
	sw.Do("}\n", nil)

	sw.Do("(*out)[*outKey] = *outVal\n", nil)
	sw.Do("}\n", nil)

	sw.Do("}\n", nil)
}

func (g *genConversion) doSlice(inType, outType *types.Type, sw *generator.SnippetWriter) {
	sw.Do("if *in == nil {\n", nil)
	sw.Do("*out = nil\n", nil)
	sw.Do("} else {\n", nil)

	sw.Do("*out = make($.|raw$, len(*in))\n", outType)
	if inType.Elem == outType.Elem && inType.Elem.Kind == types.Builtin {
		sw.Do("copy(*out, *in)\n", nil)
	} else {
		sw.Do("for i := range *in {\n", nil)
		sw.Do("in, out := &(*in)[i], &(*out)[i]\n", nil)
		g.generateFor(inType.Elem, outType.Elem, true, sw)
		sw.Do("}\n", nil)
	}

	sw.Do("}\n", nil)
}

func (g *genConversion) doStruct(inType, outType *types.Type, sw *generator.SnippetWriter) {
	for _, inMember := range inType.Members {

		if tagvals := extractTag(inMember.CommentLines); tagvals != nil && tagvals[0] == "false" {
			// This field is excluded from conversion.
			sw.Do("// INFO: in.$.inMember.Name$ opted out of conversion generation via +k8s:conversion-gen=false\n", generator.Args{
				"inMember": inMember,
			})
			continue
		}
		outMember, found := findMember(outType, append([]string{inMember.Name}, extractRenameTags(inMember.CommentLines)...)...)
		if !found {
			// This field doesn't exist in the peer.
			sw.Do("// WARNING: in.$.inMember.Name$ requires manual conversion: does not exist in peer-type\n", generator.Args{
				"inMember": inMember,
			})
			g.skippedFields[inType] = append(g.skippedFields[inType], inMember.Name)
			continue
		}

		if function, ok := g.preexists(inMember.Type, outMember.Type); ok && isDrop(function.CommentLines) {
			// This field is excluded from conversion.

			// sw.Do("// INFO: in.$.inMember.Name$ opted out of conversion generation via +k8s:conversion-fn=drop on $.function|raw$\n", generator.Args{
			// 	"inMember": inMember,
			// 	"function": function,
			// })
			continue
		}

		sw.Do("if true {\n", nil)
		sw.Do("in, out := &in.$.inMember.Name$, &out.$.outMember.Name$\n", generator.Args{
			"inMember":  inMember,
			"outMember": outMember,
		})
		g.generateFor(inMember.Type, outMember.Type, true, sw)
		sw.Do("}\n", nil)
	}
}

func (g *genConversion) doPointer(inType, outType *types.Type, sw *generator.SnippetWriter) {
	sw.Do("if *in == nil {\n", nil)
	sw.Do("*out = nil\n", nil)
	sw.Do("} else {\n", nil)
	sw.Do("*out = new($.Elem|raw$)\n", outType)
	sw.Do("in, out := *in, *out\n", nil)
	g.generateFor(inType.Elem, outType.Elem, true, sw)
	sw.Do("}\n", nil)
}

func (g *genConversion) generateFromUrlValues(inType, outType *types.Type, sw *generator.SnippetWriter) {
	args := generator.Args{
		"inType":  inType,
		"outType": outType,
		"Scope":   types.Ref(conversionPackagePath, "Scope"),
	}
	sw.Do("func auto"+nameTmpl+"(in *$.inType|raw$, out *$.outType|raw$, s $.Scope|raw$) error {\n", args)
	for _, outMember := range outType.Members {
		if tagvals := extractTag(outMember.CommentLines); tagvals != nil && tagvals[0] == "false" {
			// This field is excluded from conversion.
			sw.Do("// INFO: in."+outMember.Name+" opted out of conversion generation\n", nil)
			continue
		}
		jsonTag := reflect.StructTag(outMember.Tags).Get("json")
		index := strings.Index(jsonTag, ",")
		if index == -1 {
			index = len(jsonTag)
		}
		if index == 0 {
			memberArgs := generator.Args{
				"name": outMember.Name,
			}
			sw.Do("// WARNING: Field $.name$ does not have json tag, skipping.\n\n", memberArgs)
			continue
		}
		memberArgs := generator.Args{
			"name": outMember.Name,
			"tag":  jsonTag[:index],
		}
		sw.Do("if values, ok := map[string][]string(*in)[\"$.tag$\"]; ok && len(values) > 0 {\n", memberArgs)
		g.fromValuesEntry(inType.Underlying.Elem, outMember, sw)
		sw.Do("} else {\n", nil)
		g.setZeroValue(outMember, sw)
		sw.Do("}\n", nil)
	}
	sw.Do("return nil\n", nil)
	sw.Do("}\n\n", nil)

	if _, found := g.preexists(inType, outType); found {
		// There is a public manual Conversion method: use it.
	} else {
		// Emit a public conversion function.
		sw.Do("// "+nameTmpl+" is an autogenerated conversion function.\n", args)
		sw.Do("func "+nameTmpl+"(in *$.inType|raw$, out *$.outType|raw$, s $.Scope|raw$) error {\n", args)
		sw.Do("return auto"+nameTmpl+"(in, out, s)\n", args)
		sw.Do("}\n\n", nil)
	}
}

func (g *genConversion) fromValuesEntry(inType *types.Type, outMember types.Member, sw *generator.SnippetWriter) {
	memberArgs := generator.Args{
		"name": outMember.Name,
		"type": outMember.Type,
	}
	if function, ok := g.preexists(inType, outMember.Type); ok {
		args := memberArgs.With("function", function)
		sw.Do("if err := $.function|raw$(&values, &out.$.name$, s); err != nil {\n", args)
		sw.Do("return err\n", nil)
		sw.Do("}\n", nil)
		return
	}
	switch {
	case outMember.Type == types.String:
		sw.Do("out.$.name$ = values[0]\n", memberArgs)
	case g.useUnsafe.Equal(inType, outMember.Type):
		args := memberArgs.With("Pointer", types.Ref("unsafe", "Pointer"))
		switch inType.Kind {
		case types.Pointer:
			sw.Do("out.$.name$ = ($.type|raw$)($.Pointer|raw$(&values))\n", args)
		case types.Map, types.Slice:
			sw.Do("out.$.name$ = *(*$.type|raw$)($.Pointer|raw$(&values))\n", args)
		default:
			// TODO: Support other types to allow more auto-conversions.
			sw.Do("// FIXME: out.$.name$ is of not yet supported type and requires manual conversion\n", memberArgs)
		}
	default:
		// TODO: Support other types to allow more auto-conversions.
		sw.Do("// FIXME: out.$.name$ is of not yet supported type and requires manual conversion\n", memberArgs)
	}
}

func (g *genConversion) setZeroValue(outMember types.Member, sw *generator.SnippetWriter) {
	outMemberType := unwrapAlias(outMember.Type)
	memberArgs := generator.Args{
		"name":  outMember.Name,
		"alias": outMember.Type,
		"type":  outMemberType,
	}

	switch outMemberType.Kind {
	case types.Builtin:
		switch outMemberType {
		case types.String:
			sw.Do("out.$.name$ = \"\"\n", memberArgs)
		case types.Int64, types.Int32, types.Int16, types.Int, types.Uint64, types.Uint32, types.Uint16, types.Uint:
			sw.Do("out.$.name$ = 0\n", memberArgs)
		case types.Uintptr, types.Byte:
			sw.Do("out.$.name$ = 0\n", memberArgs)
		case types.Float64, types.Float32, types.Float:
			sw.Do("out.$.name$ = 0\n", memberArgs)
		case types.Bool:
			sw.Do("out.$.name$ = false\n", memberArgs)
		default:
			sw.Do("// FIXME: out.$.name$ is of unsupported type and requires manual conversion\n", memberArgs)
		}
	case types.Struct:
		if outMemberType == outMember.Type {
			sw.Do("out.$.name$ = $.type|raw${}\n", memberArgs)
		} else {
			sw.Do("out.$.name$ = $.alias|raw$($.type|raw${})\n", memberArgs)
		}
	case types.Map, types.Slice, types.Pointer:
		sw.Do("out.$.name$ = nil\n", memberArgs)
	case types.Alias:
		// outMemberType was already unwrapped from aliases - so that should never happen.
		sw.Do("// FIXME: unexpected error for out.$.name$\n", memberArgs)
	case types.Interface, types.Array:
		sw.Do("out.$.name$ = nil\n", memberArgs)
	default:
		sw.Do("// FIXME: out.$.name$ is of unsupported type and requires manual conversion\n", memberArgs)
	}
}

// isDirectlyAssignable mimics stdlib reflect.directlyAssignable.
func isDirectlyAssignable(inType, outType *types.Type) bool {
	return outType == inType
}

// isDirectlyConvertible mimics stdlib `reflect.convertOp(outType, inType) == reflect.cvtDirect`.
func isDirectlyConvertible(inType, outType *types.Type) bool {
	return haveIdenticalUnderlyingType(outType, inType, false)
}

// haveIndenticalType mimics stdlib reflect.haveIdenticalType.
func haveIdenticalType(outType, inType *types.Type, cmpTags bool) bool {
	if cmpTags {
		return outType == inType
	}

	if outType.Name.Name != inType.Name.Name || outType.Kind != inType.Kind || outType.Name.Package != inType.Name.Package {
		return false
	}

	return haveIdenticalUnderlyingType(outType, inType, false)
}

// haveIndenticalUnderlyingType mimics stdlib reflect.haveIndenticalUnderlyingType.
func haveIdenticalUnderlyingType(outType, inType *types.Type, cmpTags bool) bool {
	outType, inType = unwrapAlias(outType), unwrapAlias(inType)
	if outType == inType {
		return true
	}
	if inType.Kind != outType.Kind {
		return false
	}
	switch inType.Kind {
	case types.Builtin:
		return inType.Name == outType.Name
	case types.Struct:
		if len(inType.Members) != len(outType.Members) {
			return false
		}
		for i := range inType.Members {
			if !(inType.Members[i].Name == outType.Members[i].Name &&
				haveIdenticalType(inType.Members[i].Type, outType.Members[i].Type, cmpTags) &&
				inType.Members[i].Embedded == outType.Members[i].Embedded) {
				return false
			}
			if cmpTags && inType.Members[i].Tags != outType.Members[i].Tags {
				return false
			}
		}
		return true
	case types.Map:
		return haveIdenticalType(inType.Elem, outType.Elem, cmpTags) &&
			haveIdenticalType(inType.Key, outType.Key, cmpTags)
	case types.Slice:
		return haveIdenticalType(inType.Elem, outType.Elem, cmpTags)
	case types.Pointer:
		return haveIdenticalType(inType.Elem, outType.Elem, cmpTags)
	case types.Interface:
		// Only consider empty interfaces to be equal; two interfaces might have the same
		// methods but still need a run time conversion.
		return len(inType.Methods) == 0 && len(outType.Methods) == 0
	case types.Func:
		outParams := outType.Signature.Parameters
		if outType.Signature.Receiver != nil {
			outParams = append([]*types.Type{outType.Signature.Receiver}, outParams...)
		}
		inParams := inType.Signature.Parameters
		if inType.Signature.Receiver != nil {
			inParams = append([]*types.Type{inType.Signature.Receiver}, inParams...)
		}
		if len(outParams) != len(inParams) || outType.Signature.Variadic != inType.Signature.Variadic {
			return false
		}
		for i := range outParams {
			if !haveIdenticalType(outParams[i], inParams[i], cmpTags) {
				return false
			}
		}
		if len(outType.Signature.Results) != len(inType.Signature.Results) {
			return false
		}
		for i := range outType.Signature.Results {
			if !haveIdenticalType(outType.Signature.Results[i], inType.Signature.Results[i], cmpTags) {
				return false
			}
		}
		return true
	case types.Array:
		// Not implemented because k8s.io/gengo/types.Type doesn't track array length.
		return false
	case types.Chan:
		// Not implemented because k8s.io/gengo/types.Type doesn't track channel direction.
		return false
	default:
		return false
	}
}

func isSamePackage(inType, outType *types.Type) bool {
	return inType.Name.Package == outType.Name.Package
}