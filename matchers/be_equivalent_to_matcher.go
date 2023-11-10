package matchers

import (
	"github.com/onsi/gomega/format"

	"github.com/shopspring/decimal"
)

type BeEquivalentToMatcher struct {
	Expected decimal.Decimal
}

func (matcher *BeEquivalentToMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil {
		return false, ErrActualNotBeNil
	}

	if actualDecimal, ok := actual.(decimal.Decimal); ok {
		return actualDecimal.Equal(matcher.Expected), nil
	}

	if s, ok := actual.(string); ok {
		actualDecimal, err := decimal.NewFromString(s)
		if err != nil {
			return false, ErrFailedConvertStringToDecimal
		}

		return actualDecimal.Equal(matcher.Expected), nil
	}

	return false, ErrNotSupportedValue
}

func (matcher *BeEquivalentToMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be equivalent to", matcher.Expected)
}

func (matcher *BeEquivalentToMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be equivalent to", matcher.Expected)
}
