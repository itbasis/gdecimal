package matchers_test

import (
	"fmt"

	gdecimalErrors "github.com/itbasis/gdecimal/errors"
	converterTestData "github.com/itbasis/gdecimal/internal/converter/testdata"
	"github.com/itbasis/gdecimal/internal/matchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Context(
	"BeLessThanOrEqualMatcher", func() {

		Describe(
			"asserting on nil", func() {
				success, err := (&matchers.BeLessThanOrEqualMatcher{Expected: decimal.NewFromInt(1)}).Match(nil)
				Ω(success).Should(BeFalse())
				Ω(err).Should(HaveOccurred())
			},
		)

		DescribeTable(
			"passed matchers", func(expected, actual interface{}) {
				Ω((&matchers.BeLessThanOrEqualMatcher{Expected: expected}).Match(actual)).Should(BeTrue())
			},
			Entry(nil, decimal.NewFromInt(1), decimal.NewFromInt(0)),
			Entry(nil, decimal.NewFromInt(1), decimal.NewFromInt(1)),
			Entry(nil, decimal.New(1, 1), decimal.New(1, 1)),
			Entry(nil, decimal.New(11, 1), decimal.New(1, 1)),
			Entry(nil, decimal.New(1, 1), decimal.New(9, 0)),
			Entry(nil, "1", "1"),
			Entry(nil, "1", "0.9"),
			Entry(nil, decimal.NewFromInt(1), "1"),
			Entry(nil, decimal.NewFromInt(1), "0.9"),
			Entry(nil, decimal.New(1, 1), "10"),
			Entry(nil, decimal.New(1, 0), decimal.RequireFromString("0.9")),
			Entry(nil, decimal.New(1, 1), decimal.RequireFromString("9")),
			Entry(nil, decimal.New(11, -1), decimal.RequireFromString("1.0")),
			Entry(nil, decimal.NewFromInt(1), "0.9"),
			Entry(nil, decimal.New(1, 1), "9"),
			Entry(nil, decimal.New(1, 1), "10"),
			Entry(nil, decimal.New(11, 1), "109"),
			Entry(nil, decimal.New(11, 1), "110"),
			Entry(nil, decimal.New(11, -1), "1.1"),
			Entry(nil, decimal.New(11, -1), "1.0"),
		)

		DescribeTable(
			"passed matchers (Not)", func(expected, actual interface{}) {
				Ω((&matchers.BeLessThanOrEqualMatcher{Expected: expected}).Match(actual)).ShouldNot(BeTrue())
			},
			Entry(nil, decimal.New(9, 0), decimal.New(1, 1)),
			Entry(nil, "10", "11"),
			Entry(nil, decimal.New(1, 1), "11"),
			Entry(nil, decimal.New(1, -1), "1"),
		)

		DescribeTable(
			"should use the matcher's failure messages",
			func(expected, actual interface{}, expectedValue string) {
				matcher := &matchers.BeLessThanOrEqualMatcher{Expected: expected}
				Ω(matcher.FailureMessage(actual)).Should(
					Equal(
						fmt.Sprintf(
							"Expected\n    <string>: %s\nto be less than or equal to\n    <decimal.Decimal>: {\n        value: %s",
							actual,
							expectedValue,
						),
					),
				)

				Ω(matcher.NegatedFailureMessage(actual)).Should(
					Equal(
						fmt.Sprintf(
							"Expected\n    <string>: %s\nnot to be less than or equal to\n    <decimal.Decimal>: {\n        value: %s",
							actual,
							expectedValue,
						),
					),
				)
			},
			Entry(
				nil,
				decimal.NewFromInt(10),
				"1",
				"{neg: false, abs: [10]},\n        exp: 0,\n    }",
			),
			Entry(
				nil,
				decimal.New(11, -1),
				"11",
				"{neg: false, abs: [11]},\n        exp: -1,\n    }",
			),
		)

		Describe(
			"asserting failed convert", func() {
				DescribeTable(
					"actual on failure string", func(value string) {
						success, err := (&matchers.BeLessThanOrEqualMatcher{Expected: decimal.Decimal{}}).Match(value)
						Ω(success).To(BeFalse())
						Ω(err).Should(MatchError(gdecimalErrors.ErrFailedConvertStringToDecimal))
					},
					converterTestData.AssertingOnStringFailureEntries,
				)
				DescribeTable(
					"expected on failure string", func(value string) {
						success, err := (&matchers.BeLessThanOrEqualMatcher{Expected: value}).Match(decimal.Decimal{})
						Ω(success).To(BeFalse())
						Ω(err).Should(MatchError(gdecimalErrors.ErrFailedConvertStringToDecimal))
					},
					converterTestData.AssertingOnStringFailureEntries,
				)
			},
		)

	},
)
