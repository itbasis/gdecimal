package gdecimal

import (
	"github.com/itbasis/gdecimal/matchers"
	"github.com/shopspring/decimal"
)

func BeDecimalEquivalentTo(expect decimal.Decimal) *matchers.BeEquivalentToMatcher {
	return &matchers.BeEquivalentToMatcher{Expected: expect}
}
