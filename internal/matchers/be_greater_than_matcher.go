package matchers

import (
	"github.com/itbasis/gdecimal/internal/converter"
	"github.com/onsi/gomega/format"
)

type BeGreaterThanMatcher struct {
	Expected interface{}
}

func (matcher *BeGreaterThanMatcher) Match(actual interface{}) (bool, error) {
	actualDecimal, errActual := converter.ToDecimal(actual)
	if errActual != nil {
		return false, errActual
	}

	expectedDecimal, errExpected := converter.ToDecimal(matcher.Expected)
	if errExpected != nil {
		return false, errExpected
	}

	return actualDecimal.GreaterThan(expectedDecimal), nil
}

func (matcher *BeGreaterThanMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to be greater than to", matcher.Expected)
}

func (matcher *BeGreaterThanMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to be greater than to", matcher.Expected)
}
