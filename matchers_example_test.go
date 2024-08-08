package gdecimal_test

import (
	. "github.com/itbasis/gdecimal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe(
	"sample", func() {
		defer GinkgoRecover()

		expected := decimal.NewFromInt(10)

		// as Decimal
		Expect(decimal.NewFromInt(10)).To(BeDecimalEquivalentTo(expected))
		// as string
		Expect("10").To(BeDecimalEquivalentTo(expected))

		// expect as string
		Expect(decimal.NewFromInt(10)).To(BeDecimalEquivalentTo("10"))
	},
)
