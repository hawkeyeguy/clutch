// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authn/v1/authn.proto

package authnv1

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

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RedirectUrl

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResponseMultiError, or
// nil if none found.
func (m *LoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Return.(type) {
	case *LoginResponse_AuthUrl:
		if v == nil {
			err := LoginResponseValidationError{
				field:  "Return",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for AuthUrl
	case *LoginResponse_Token_:
		if v == nil {
			err := LoginResponseValidationError{
				field:  "Return",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetToken()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LoginResponseValidationError{
						field:  "Token",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LoginResponseValidationError{
						field:  "Token",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LoginResponseValidationError{
					field:  "Token",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return LoginResponseMultiError(errors)
	}

	return nil
}

// LoginResponseMultiError is an error wrapping multiple validation errors
// returned by LoginResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponseMultiError) AllErrors() []error { return m }

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on CallbackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CallbackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallbackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CallbackRequestMultiError, or nil if none found.
func (m *CallbackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CallbackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for State

	// no validation rules for Error

	// no validation rules for ErrorDescription

	if len(errors) > 0 {
		return CallbackRequestMultiError(errors)
	}

	return nil
}

// CallbackRequestMultiError is an error wrapping multiple validation errors
// returned by CallbackRequest.ValidateAll() if the designated constraints
// aren't met.
type CallbackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallbackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallbackRequestMultiError) AllErrors() []error { return m }

// CallbackRequestValidationError is the validation error returned by
// CallbackRequest.Validate if the designated constraints aren't met.
type CallbackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallbackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallbackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallbackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallbackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallbackRequestValidationError) ErrorName() string { return "CallbackRequestValidationError" }

// Error satisfies the builtin error interface
func (e CallbackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallbackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallbackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallbackRequestValidationError{}

// Validate checks the field values on CallbackResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CallbackResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CallbackResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CallbackResponseMultiError, or nil if none found.
func (m *CallbackResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CallbackResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return CallbackResponseMultiError(errors)
	}

	return nil
}

// CallbackResponseMultiError is an error wrapping multiple validation errors
// returned by CallbackResponse.ValidateAll() if the designated constraints
// aren't met.
type CallbackResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CallbackResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CallbackResponseMultiError) AllErrors() []error { return m }

// CallbackResponseValidationError is the validation error returned by
// CallbackResponse.Validate if the designated constraints aren't met.
type CallbackResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallbackResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallbackResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallbackResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallbackResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallbackResponseValidationError) ErrorName() string { return "CallbackResponseValidationError" }

// Error satisfies the builtin error interface
func (e CallbackResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallbackResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallbackResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallbackResponseValidationError{}

// Validate checks the field values on CreateTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTokenRequestMultiError, or nil if none found.
func (m *CreateTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetSubject()) < 1 {
		err := CreateTokenRequestValidationError{
			field:  "Subject",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetExpiry()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTokenRequestValidationError{
					field:  "Expiry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTokenRequestValidationError{
					field:  "Expiry",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiry()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTokenRequestValidationError{
				field:  "Expiry",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := _CreateTokenRequest_TokenType_NotInLookup[m.GetTokenType()]; ok {
		err := CreateTokenRequestValidationError{
			field:  "TokenType",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := CreateTokenRequest_TokenType_name[int32(m.GetTokenType())]; !ok {
		err := CreateTokenRequestValidationError{
			field:  "TokenType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateTokenRequestMultiError(errors)
	}

	return nil
}

// CreateTokenRequestMultiError is an error wrapping multiple validation errors
// returned by CreateTokenRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenRequestMultiError) AllErrors() []error { return m }

// CreateTokenRequestValidationError is the validation error returned by
// CreateTokenRequest.Validate if the designated constraints aren't met.
type CreateTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenRequestValidationError) ErrorName() string {
	return "CreateTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenRequestValidationError{}

var _CreateTokenRequest_TokenType_NotInLookup = map[CreateTokenRequest_TokenType]struct{}{
	0: {},
}

// Validate checks the field values on CreateTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTokenResponseMultiError, or nil if none found.
func (m *CreateTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return CreateTokenResponseMultiError(errors)
	}

	return nil
}

// CreateTokenResponseMultiError is an error wrapping multiple validation
// errors returned by CreateTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTokenResponseMultiError) AllErrors() []error { return m }

// CreateTokenResponseValidationError is the validation error returned by
// CreateTokenResponse.Validate if the designated constraints aren't met.
type CreateTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTokenResponseValidationError) ErrorName() string {
	return "CreateTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTokenResponseValidationError{}

// Validate checks the field values on LoginResponse_Token with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LoginResponse_Token) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse_Token with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginResponse_TokenMultiError, or nil if none found.
func (m *LoginResponse_Token) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse_Token) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return LoginResponse_TokenMultiError(errors)
	}

	return nil
}

// LoginResponse_TokenMultiError is an error wrapping multiple validation
// errors returned by LoginResponse_Token.ValidateAll() if the designated
// constraints aren't met.
type LoginResponse_TokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponse_TokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponse_TokenMultiError) AllErrors() []error { return m }

// LoginResponse_TokenValidationError is the validation error returned by
// LoginResponse_Token.Validate if the designated constraints aren't met.
type LoginResponse_TokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponse_TokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponse_TokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponse_TokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponse_TokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponse_TokenValidationError) ErrorName() string {
	return "LoginResponse_TokenValidationError"
}

// Error satisfies the builtin error interface
func (e LoginResponse_TokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse_Token.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponse_TokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponse_TokenValidationError{}
