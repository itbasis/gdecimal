package errors

import "errors"

var (
	ErrNotBeNil                     = errors.New("must not be nil")
	ErrFailedConvertStringToDecimal = errors.New("DecimalBeEquivalentToMatcher failed to convert string to Decimal")
	ErrNotSupportedValue            = errors.New("DecimalBeEquivalentToMatcher matcher expects a Decimal or string")
)
