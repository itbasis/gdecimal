package matchers_test

import (
	"fmt"

	. "github.com/itbasis/gdecimal" //nolint:revive
	"github.com/itbasis/gdecimal/matchers"
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
	"github.com/shopspring/decimal"
)

var _ = Describe(
	"asserting on nil", func() {
		success, err := (&matchers.BeEquivalentToMatcher{Expected: decimal.NewFromInt(1)}).Match(nil)
		Ω(success).Should(BeFalse())
		Ω(err).Should(HaveOccurred())
	},
)

var _ = Describe(
	"asserting on Decimal", func() {
		DescribeTable(
			"passed matchers", func(expected decimal.Decimal, actual interface{}) {
				Ω((&matchers.BeEquivalentToMatcher{Expected: expected}).Match(actual)).Should(BeTrue())
			},
			Entry(nil, decimal.NewFromInt(1), decimal.NewFromInt(1)),
			Entry(nil, decimal.New(1, 1), decimal.New(1, 1)),
			Entry(nil, decimal.New(1, 0), decimal.RequireFromString("1")),
			Entry(nil, decimal.New(1, 1), decimal.RequireFromString("10")),
			Entry(nil, decimal.New(11, -1), decimal.RequireFromString("1.1")),
		)

		DescribeTable(
			"passed matchers (Not)", func(expected decimal.Decimal, actual interface{}) {
				Ω((&matchers.BeEquivalentToMatcher{Expected: expected}).Match(actual)).ShouldNot(BeTrue())
			},
			Entry(nil, decimal.NewFromInt(1), decimal.NewFromInt(2)),
			Entry(nil, decimal.New(1, 1), decimal.New(1, 2)),
			Entry(nil, decimal.New(1, 1), decimal.New(2, 1)),
		)

		DescribeTable(
			"should use the matcher's failure message (as Decimal)",
			func(expected decimal.Decimal, actual interface{}, expectedValue, actualValue string) {
				expectedErr := fmt.Sprintf(
					"Expected\n    <decimal.Decimal>: {\n        value: %s\nto be equivalent to\n    <decimal.Decimal>: {\n        value: %s",
					actualValue,
					expectedValue,
				)

				err := InterceptGomegaFailure(
					func() {
						Ω(actual).Should(BeDecimalEquivalentTo(expected))
					},
				)
				Ω(err.Error()).Should(Equal(expectedErr))

				matcher := &matchers.BeEquivalentToMatcher{Expected: expected}
				Ω(matcher.Match(actual)).Should(BeFalse())
				Ω(matcher.FailureMessage(actual)).Should(Equal(expectedErr))
			},
			Entry(
				nil,
				decimal.NewFromInt(10),
				decimal.NewFromInt(1),
				"{neg: false, abs: [10]},\n        exp: 0,\n    }",
				"{neg: false, abs: [1]},\n        exp: 0,\n    }",
			),
			Entry(
				nil,
				decimal.New(11, -1),
				decimal.NewFromInt(11),
				"{neg: false, abs: [11]},\n        exp: -1,\n    }",
				"{neg: false, abs: [11]},\n        exp: 0,\n    }",
			),
		)
	},
)

var _ = Describe(
	"asserting on string", func() {
		DescribeTable(
			"passed matchers", func(expected decimal.Decimal, actual interface{}) {
				Ω((&matchers.BeEquivalentToMatcher{Expected: expected}).Match(actual)).Should(BeTrue())
			},
			Entry(nil, decimal.NewFromInt(1), "1"),
			Entry(nil, decimal.New(1, 1), "10"),
			Entry(nil, decimal.New(11, 1), "110"),
			Entry(nil, decimal.New(11, -1), "1.1"),
		)

		DescribeTable(
			"passed matchers (Not)", func(expected decimal.Decimal, actual interface{}) {
				Ω((&matchers.BeEquivalentToMatcher{Expected: expected}).Match(actual)).ShouldNot(BeTrue())
			},
			Entry(nil, decimal.NewFromInt(1), "2"),
			Entry(nil, decimal.New(1, 1), "1"),
			Entry(nil, decimal.New(1, 1), "20"),
			Entry(nil, decimal.New(1, 1), "1.1"),
		)

		DescribeTable(
			"should use the matcher's failure message",
			func(expected decimal.Decimal, actual interface{}, expectedValue string) {
				expectedErr := fmt.Sprintf(
					"Expected\n    <string>: %s\nto be equivalent to\n    <decimal.Decimal>: {\n        value: %s",
					actual,
					expectedValue,
				)

				err := InterceptGomegaFailure(
					func() {
						Ω(actual).Should(BeDecimalEquivalentTo(expected))
					},
				)
				Ω(err.Error()).Should(Equal(expectedErr))

				matcher := &matchers.BeEquivalentToMatcher{Expected: expected}
				Ω(matcher.Match(actual)).Should(BeFalse())
				Ω(matcher.FailureMessage(actual)).Should(Equal(expectedErr))
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
	},
)
