package testdata

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	"github.com/shopspring/decimal"
)

var AssertingOnStringPassedEntries = []TableEntry{
	Entry(nil, "1", decimal.NewFromInt(1)),
	Entry(nil, "-1", decimal.NewFromInt(-1)),
	Entry(nil, "1.0", decimal.New(10, -1)),
	Entry(nil, "01", decimal.NewFromInt(1)),
}
var AssertingOnStringFailureEntries = []TableEntry{
	Entry(nil, ""),
	Entry(nil, " 2"),
	Entry(nil, "2 "),
	Entry(nil, "2,0"),
}
var AssertingOnAnotherEntries = []TableEntry{
	Entry(nil, 1.0),
	Entry(nil, -1.0),
}
