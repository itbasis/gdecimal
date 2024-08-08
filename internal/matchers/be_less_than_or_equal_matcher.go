package matchers

import (
	"github.com/itbasis/gdecimal/internal/converter"
	"github.com/onsi/gomega/format"
)

type BeLessThanOrEqualMatcher struct {
	Expected interface{}
}

func (matcher *BeLessThanOrEqualMatcher) Match(actual interface{}) (bool, error) {
	actualDecimal, errActual := converter.ToDecimal(actual)
	if errActual != nil {
		return false, errActual
	}

	expectedDecimal, errExpected := converter.ToDecimal(matcher.Expected)
	if errExpected != nil {
		return false, errExpected
	}

	return actualDecimal.LessThanOrEqual(expectedDecimal), nil
}

func (matcher *BeLessThanOrEqualMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to be less than or equal to", matcher.Expected)
}

func (matcher *BeLessThanOrEqualMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to be less than or equal to", matcher.Expected)
}
