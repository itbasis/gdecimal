package matchers

import (
	"github.com/itbasis/gdecimal/internal/converter"
	"github.com/onsi/gomega/format"
)

type BeLessThanMatcher struct {
	Expected interface{}
}

func (matcher *BeLessThanMatcher) Match(actual interface{}) (success bool, err error) {
	actualDecimal, errActual := converter.ToDecimal(actual)
	if errActual != nil {
		return false, errActual //nolint:wrapcheck
	}

	expectedDecimal, errExpected := converter.ToDecimal(matcher.Expected)
	if errExpected != nil {
		return false, errExpected //nolint:wrapcheck
	}

	return actualDecimal.LessThan(expectedDecimal), nil
}

func (matcher *BeLessThanMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to be less than to", matcher.Expected)
}

func (matcher *BeLessThanMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to be less than to", matcher.Expected)
}
