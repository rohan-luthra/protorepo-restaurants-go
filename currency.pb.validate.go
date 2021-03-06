// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: currency.proto

package restaurant

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

// define the regex for a UUID once up-front
var _currency_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Currency with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Currency) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Symbol

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CurrencyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CurrencyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CurrencyValidationError is the validation error returned by
// Currency.Validate if the designated constraints aren't met.
type CurrencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrencyValidationError) ErrorName() string { return "CurrencyValidationError" }

// Error satisfies the builtin error interface
func (e CurrencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrencyValidationError{}
