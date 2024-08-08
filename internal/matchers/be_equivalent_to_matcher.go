package matchers

import (
	"github.com/itbasis/gdecimal/internal/converter"
	"github.com/onsi/gomega/format"
)

type BeEquivalentToMatcher struct {
	Expected interface{}
}

func (matcher *BeEquivalentToMatcher) Match(actual interface{}) (bool, error) {
	actualDecimal, errActual := converter.ToDecimal(actual)
	if errActual != nil {
		return false, errActual
	}

	expectedDecimal, errExpected := converter.ToDecimal(matcher.Expected)
	if errExpected != nil {
		return false, errExpected
	}

	return actualDecimal.Equal(expectedDecimal), nil
}

func (matcher *BeEquivalentToMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to be equivalent to", matcher.Expected)
}

func (matcher *BeEquivalentToMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to be equivalent to", matcher.Expected)
}
