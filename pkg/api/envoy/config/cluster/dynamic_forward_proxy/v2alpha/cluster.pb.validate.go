// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/cluster/dynamic_forward_proxy/v2alpha/cluster.proto

package envoy_config_cluster_dynamic_forward_proxy_v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ClusterConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClusterConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDnsCacheConfig() == nil {
		return ClusterConfigValidationError{
			field:  "DnsCacheConfig",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDnsCacheConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterConfigValidationError{
				field:  "DnsCacheConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ClusterConfigValidationError is the validation error returned by
// ClusterConfig.Validate if the designated constraints aren't met.
type ClusterConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterConfigValidationError) ErrorName() string { return "ClusterConfigValidationError" }

// Error satisfies the builtin error interface
func (e ClusterConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterConfigValidationError{}
