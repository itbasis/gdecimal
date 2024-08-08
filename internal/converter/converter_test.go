package converter_test

import (
	gdecimalErrors "github.com/itbasis/gdecimal/errors"
	"github.com/itbasis/gdecimal/internal/converter"
	converterTestData "github.com/itbasis/gdecimal/internal/converter/testdata"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe(
	"asserting on nil", func() {
		Ω(converter.ToDecimal(nil)).Error().To(MatchError(gdecimalErrors.ErrNotBeNil))
	},
)

var _ = Describe(
	"asserting on Decimal", func() {
		Ω(converter.ToDecimal(decimal.Decimal{})).To(Equal(decimal.Decimal{}))
	},
)

var _ = Describe(
	"asserting on string", func() {
		DescribeTable(
			"passed", func(value string, expect decimal.Decimal) {
				Ω(converter.ToDecimal(value)).To(Equal(expect))
			},
			converterTestData.AssertingOnStringPassedEntries,
		)

		DescribeTable(
			"failure", func(value string) {
				Ω(converter.ToDecimal(value)).Error().To(MatchError(gdecimalErrors.ErrFailedConvertStringToDecimal))
			},
			converterTestData.AssertingOnStringFailureEntries,
		)
	},
)

var _ = DescribeTable(
	"asserting on another", func(value interface{}) {
		Ω(converter.ToDecimal(value)).Error().To(MatchError(gdecimalErrors.ErrNotSupportedValue))
	},
	converterTestData.AssertingOnAnotherEntries,
)
