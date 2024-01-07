package matchers

import (
	"github.com/itbasis/gdecimal/internal/converter"
	"github.com/onsi/gomega/format"
)

type BeLessThanOrEqualMatcher struct {
	Expected interface{}
}

func (matcher *BeLessThanOrEqualMatcher) Match(actual interface{}) (success bool, err error) {
	actualDecimal, errActual := converter.ToDecimal(actual)
	if errActual != nil {
		return false, errActual //nolint:wrapcheck
	}

	expectedDecimal, errExpected := converter.ToDecimal(matcher.Expected)
	if errExpected != nil {
		return false, errExpected //nolint:wrapcheck
	}

	return actualDecimal.LessThanOrEqual(expectedDecimal), nil
}

func (matcher *BeLessThanOrEqualMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be less than or equal to", matcher.Expected)
}

func (matcher *BeLessThanOrEqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be less than or equal to", matcher.Expected)
}
