// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: feedback/v1/feedback.proto

package feedbackv1

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
)

// Validate checks the field values on GetQuestionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetQuestionsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetOrigins()) < 1 {
		return GetQuestionsRequestValidationError{
			field:  "Origins",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetOrigins() {
		_, _ = idx, item

		if _, ok := _GetQuestionsRequest_Origins_NotInLookup[item]; ok {
			return GetQuestionsRequestValidationError{
				field:  fmt.Sprintf("Origins[%v]", idx),
				reason: "value must not be in list [0]",
			}
		}

		if _, ok := Origin_name[int32(item)]; !ok {
			return GetQuestionsRequestValidationError{
				field:  fmt.Sprintf("Origins[%v]", idx),
				reason: "value must be one of the defined enum values",
			}
		}

	}

	if len(m.GetUser()) < 1 {
		return GetQuestionsRequestValidationError{
			field:  "User",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// GetQuestionsRequestValidationError is the validation error returned by
// GetQuestionsRequest.Validate if the designated constraints aren't met.
type GetQuestionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetQuestionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetQuestionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetQuestionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetQuestionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetQuestionsRequestValidationError) ErrorName() string {
	return "GetQuestionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetQuestionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetQuestionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetQuestionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetQuestionsRequestValidationError{}

var _GetQuestionsRequest_Origins_NotInLookup = map[Origin]struct{}{
	0: {},
}

// Validate checks the field values on RatingOptions with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RatingOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for One

	// no validation rules for Two

	// no validation rules for Three

	return nil
}

// RatingOptionsValidationError is the validation error returned by
// RatingOptions.Validate if the designated constraints aren't met.
type RatingOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RatingOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RatingOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RatingOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RatingOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RatingOptionsValidationError) ErrorName() string { return "RatingOptionsValidationError" }

// Error satisfies the builtin error interface
func (e RatingOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRatingOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RatingOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RatingOptionsValidationError{}

// Validate checks the field values on OriginQuestion with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OriginQuestion) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Question

	// no validation rules for FreeformQuestion

	if v, ok := interface{}(m.GetRatingOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OriginQuestionValidationError{
				field:  "RatingOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Show

	return nil
}

// OriginQuestionValidationError is the validation error returned by
// OriginQuestion.Validate if the designated constraints aren't met.
type OriginQuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OriginQuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OriginQuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OriginQuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OriginQuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OriginQuestionValidationError) ErrorName() string { return "OriginQuestionValidationError" }

// Error satisfies the builtin error interface
func (e OriginQuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOriginQuestion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OriginQuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OriginQuestionValidationError{}

// Validate checks the field values on GetQuestionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetQuestionsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetOriginQuestion() {
		_ = val

		// no validation rules for OriginQuestion[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetQuestionsResponseValidationError{
					field:  fmt.Sprintf("OriginQuestion[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetQuestionsResponseValidationError is the validation error returned by
// GetQuestionsResponse.Validate if the designated constraints aren't met.
type GetQuestionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetQuestionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetQuestionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetQuestionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetQuestionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetQuestionsResponseValidationError) ErrorName() string {
	return "GetQuestionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetQuestionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetQuestionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetQuestionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetQuestionsResponseValidationError{}