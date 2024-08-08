package matchers

import (
	"github.com/itbasis/gdecimal/internal/converter"
	"github.com/onsi/gomega/format"
)

type BeGreaterThanOrEqualMatcher struct {
	Expected interface{}
}

func (matcher *BeGreaterThanOrEqualMatcher) Match(actual interface{}) (bool, error) {
	actualDecimal, errActual := converter.ToDecimal(actual)
	if errActual != nil {
		return false, errActual
	}

	expectedDecimal, errExpected := converter.ToDecimal(matcher.Expected)
	if errExpected != nil {
		return false, errExpected
	}

	return actualDecimal.GreaterThanOrEqual(expectedDecimal), nil
}

func (matcher *BeGreaterThanOrEqualMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to be greater than or equal to", matcher.Expected)
}

func (matcher *BeGreaterThanOrEqualMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to be greater than or equal to", matcher.Expected)
}
