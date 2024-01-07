package gdecimal

import (
	"github.com/itbasis/gdecimal/internal/matchers"
)

func BeDecimalEquivalentTo(expect interface{}) *matchers.BeEquivalentToMatcher {
	return &matchers.BeEquivalentToMatcher{Expected: expect}
}

func BeDecimalGreaterThan(expect interface{}) *matchers.BeGreaterThanMatcher {
	return &matchers.BeGreaterThanMatcher{Expected: expect}
}

func BeDecimalGreaterThanOrEqual(expect interface{}) *matchers.BeGreaterThanOrEqualMatcher {
	return &matchers.BeGreaterThanOrEqualMatcher{Expected: expect}
}

func BeDecimalLessThan(expect interface{}) *matchers.BeLessThanMatcher {
	return &matchers.BeLessThanMatcher{Expected: expect}
}

func BeDecimalLessThanOrEqual(expect interface{}) *matchers.BeLessThanOrEqualMatcher {
	return &matchers.BeLessThanOrEqualMatcher{Expected: expect}
}
