// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/ratelimit/v3/ratelimit.proto

package envoy_extensions_common_ratelimit_v3

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

	v3 "github.com/datawire/ambassador/pkg/api/envoy/type/v3"
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

	_ = v3.RateLimitUnit(0)
)

// Validate checks the field values on RateLimitDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitDescriptor) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetEntries()) < 1 {
		return RateLimitDescriptorValidationError{
			field:  "Entries",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitDescriptorValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitDescriptorValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RateLimitDescriptorValidationError is the validation error returned by
// RateLimitDescriptor.Validate if the designated constraints aren't met.
type RateLimitDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptorValidationError) ErrorName() string {
	return "RateLimitDescriptorValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptorValidationError{}

// Validate checks the field values on RateLimitDescriptor_Entry with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitDescriptor_Entry) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		return RateLimitDescriptor_EntryValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		return RateLimitDescriptor_EntryValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// RateLimitDescriptor_EntryValidationError is the validation error returned by
// RateLimitDescriptor_Entry.Validate if the designated constraints aren't met.
type RateLimitDescriptor_EntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptor_EntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptor_EntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptor_EntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptor_EntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptor_EntryValidationError) ErrorName() string {
	return "RateLimitDescriptor_EntryValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptor_EntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor_Entry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptor_EntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptor_EntryValidationError{}

// Validate checks the field values on RateLimitDescriptor_RateLimitOverride
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *RateLimitDescriptor_RateLimitOverride) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequestsPerUnit

	if _, ok := v3.RateLimitUnit_name[int32(m.GetUnit())]; !ok {
		return RateLimitDescriptor_RateLimitOverrideValidationError{
			field:  "Unit",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// RateLimitDescriptor_RateLimitOverrideValidationError is the validation error
// returned by RateLimitDescriptor_RateLimitOverride.Validate if the
// designated constraints aren't met.
type RateLimitDescriptor_RateLimitOverrideValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) ErrorName() string {
	return "RateLimitDescriptor_RateLimitOverrideValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor_RateLimitOverride.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptor_RateLimitOverrideValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptor_RateLimitOverrideValidationError{}
