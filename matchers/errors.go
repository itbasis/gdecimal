package matchers

import "errors"

var (
	ErrActualNotBeNil               = errors.New("actual must not be nil")
	ErrFailedConvertStringToDecimal = errors.New("DecimalBeEquivalentToMatcher failed to convert string to Decimal")
	ErrNotSupportedValue            = errors.New("DecimalBeEquivalentToMatcher matcher expects a Decimal or string")
)
