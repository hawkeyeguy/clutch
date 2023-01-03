// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/experimentation/v1/properties.proto

package experimentationv1

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

// Validate checks the field values on PropertiesList with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PropertiesList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PropertiesList with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PropertiesListMultiError,
// or nil if none found.
func (m *PropertiesList) ValidateAll() error {
	return m.validate(true)
}

func (m *PropertiesList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PropertiesListValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PropertiesListValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PropertiesListValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PropertiesListMultiError(errors)
	}

	return nil
}

// PropertiesListMultiError is an error wrapping multiple validation errors
// returned by PropertiesList.ValidateAll() if the designated constraints
// aren't met.
type PropertiesListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PropertiesListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PropertiesListMultiError) AllErrors() []error { return m }

// PropertiesListValidationError is the validation error returned by
// PropertiesList.Validate if the designated constraints aren't met.
type PropertiesListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertiesListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertiesListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertiesListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertiesListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertiesListValidationError) ErrorName() string { return "PropertiesListValidationError" }

// Error satisfies the builtin error interface
func (e PropertiesListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPropertiesList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertiesListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertiesListValidationError{}

// Validate checks the field values on PropertiesMap with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PropertiesMap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PropertiesMap with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PropertiesMapMultiError, or
// nil if none found.
func (m *PropertiesMap) ValidateAll() error {
	return m.validate(true)
}

func (m *PropertiesMap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetItems()))
		i := 0
		for key := range m.GetItems() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetItems()[key]
			_ = val

			// no validation rules for Items[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, PropertiesMapValidationError{
							field:  fmt.Sprintf("Items[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, PropertiesMapValidationError{
							field:  fmt.Sprintf("Items[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return PropertiesMapValidationError{
						field:  fmt.Sprintf("Items[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return PropertiesMapMultiError(errors)
	}

	return nil
}

// PropertiesMapMultiError is an error wrapping multiple validation errors
// returned by PropertiesMap.ValidateAll() if the designated constraints
// aren't met.
type PropertiesMapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PropertiesMapMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PropertiesMapMultiError) AllErrors() []error { return m }

// PropertiesMapValidationError is the validation error returned by
// PropertiesMap.Validate if the designated constraints aren't met.
type PropertiesMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertiesMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertiesMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertiesMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertiesMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertiesMapValidationError) ErrorName() string { return "PropertiesMapValidationError" }

// Error satisfies the builtin error interface
func (e PropertiesMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPropertiesMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertiesMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertiesMapValidationError{}

// Validate checks the field values on Property with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Property) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Property with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PropertyMultiError, or nil
// if none found.
func (m *Property) ValidateAll() error {
	return m.validate(true)
}

func (m *Property) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Label

	if all {
		switch v := interface{}(m.GetDisplayValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "DisplayValue",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PropertyValidationError{
					field:  "DisplayValue",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDisplayValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "DisplayValue",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.Value.(type) {
	case *Property_DateValue:
		if v == nil {
			err := PropertyValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetDateValue()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PropertyValidationError{
						field:  "DateValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PropertyValidationError{
						field:  "DateValue",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetDateValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PropertyValidationError{
					field:  "DateValue",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Property_StringValue:
		if v == nil {
			err := PropertyValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for StringValue
	case *Property_IntValue:
		if v == nil {
			err := PropertyValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for IntValue
	case *Property_UrlValue:
		if v == nil {
			err := PropertyValidationError{
				field:  "Value",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for UrlValue
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return PropertyMultiError(errors)
	}

	return nil
}

// PropertyMultiError is an error wrapping multiple validation errors returned
// by Property.ValidateAll() if the designated constraints aren't met.
type PropertyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PropertyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PropertyMultiError) AllErrors() []error { return m }

// PropertyValidationError is the validation error returned by
// Property.Validate if the designated constraints aren't met.
type PropertyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertyValidationError) ErrorName() string { return "PropertyValidationError" }

// Error satisfies the builtin error interface
func (e PropertyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProperty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertyValidationError{}
