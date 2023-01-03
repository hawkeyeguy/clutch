// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: featureflag/v1/featureflag.proto

package featureflagv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetFlagsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetFlagsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetFlagsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetFlagsRequestMultiError, or nil if none found.
func (m *GetFlagsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetFlagsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetFlagsRequestMultiError(errors)
	}

	return nil
}

// GetFlagsRequestMultiError is an error wrapping multiple validation errors
// returned by GetFlagsRequest.ValidateAll() if the designated constraints
// aren't met.
type GetFlagsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetFlagsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetFlagsRequestMultiError) AllErrors() []error { return m }

// GetFlagsRequestValidationError is the validation error returned by
// GetFlagsRequest.Validate if the designated constraints aren't met.
type GetFlagsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetFlagsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetFlagsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetFlagsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetFlagsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetFlagsRequestValidationError) ErrorName() string { return "GetFlagsRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetFlagsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetFlagsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetFlagsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetFlagsRequestValidationError{}

// Validate checks the field values on Flag with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Flag) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Flag with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FlagMultiError, or nil if none found.
func (m *Flag) ValidateAll() error {
	return m.validate(true)
}

func (m *Flag) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Type.(type) {
	case *Flag_BooleanValue:
		if v == nil {
			err := FlagValidationError{
				field:  "Type",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for BooleanValue
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return FlagMultiError(errors)
	}

	return nil
}

// FlagMultiError is an error wrapping multiple validation errors returned by
// Flag.ValidateAll() if the designated constraints aren't met.
type FlagMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FlagMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FlagMultiError) AllErrors() []error { return m }

// FlagValidationError is the validation error returned by Flag.Validate if the
// designated constraints aren't met.
type FlagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FlagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FlagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FlagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FlagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FlagValidationError) ErrorName() string { return "FlagValidationError" }

// Error satisfies the builtin error interface
func (e FlagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFlag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FlagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FlagValidationError{}

// Validate checks the field values on GetFlagsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetFlagsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetFlagsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetFlagsResponseMultiError, or nil if none found.
func (m *GetFlagsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetFlagsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetFlags()))
		i := 0
		for key := range m.GetFlags() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetFlags()[key]
			_ = val

			// no validation rules for Flags[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, GetFlagsResponseValidationError{
							field:  fmt.Sprintf("Flags[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, GetFlagsResponseValidationError{
							field:  fmt.Sprintf("Flags[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return GetFlagsResponseValidationError{
						field:  fmt.Sprintf("Flags[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return GetFlagsResponseMultiError(errors)
	}

	return nil
}

// GetFlagsResponseMultiError is an error wrapping multiple validation errors
// returned by GetFlagsResponse.ValidateAll() if the designated constraints
// aren't met.
type GetFlagsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetFlagsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetFlagsResponseMultiError) AllErrors() []error { return m }

// GetFlagsResponseValidationError is the validation error returned by
// GetFlagsResponse.Validate if the designated constraints aren't met.
type GetFlagsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetFlagsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetFlagsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetFlagsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetFlagsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetFlagsResponseValidationError) ErrorName() string { return "GetFlagsResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetFlagsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetFlagsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetFlagsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetFlagsResponseValidationError{}
