package gdecimal_test

import (
	"github.com/itbasis/gdecimal"
	"github.com/itbasis/gdecimal/errors"
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
	"github.com/shopspring/decimal"
)

var _ = Describe(
	"asserting on nil", func() {
		It(
			"actual is nil", func() {
				Ω(
					InterceptGomegaFailure(
						func() {
							Ω(nil).To(gdecimal.BeDecimalEquivalentTo(decimal.Decimal{}))
						},
					),
				).Error().To(Equal(errors.ErrNotBeNil))
			},
		)

		It(
			"expected is nil", func() {
				Ω(
					InterceptGomegaFailure(
						func() {
							Ω(decimal.Decimal{}).To(gdecimal.BeDecimalEquivalentTo(nil))
						},
					),
				).Error().To(Equal(errors.ErrNotBeNil))
			},
		)
	},
)
