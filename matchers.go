package gdecimal

import (
	"github.com/itbasis/gdecimal/internal/matchers"
)

func BeDecimalEquivalentTo(expect interface{}) *matchers.BeEquivalentToMatcher {
	return &matchers.BeEquivalentToMatcher{Expected: expect}
}
