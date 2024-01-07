package gdecimal_test

import (
	"github.com/itbasis/gdecimal"
	"github.com/itbasis/gdecimal/errors"
	. "github.com/onsi/ginkgo/v2" //nolint:revive
	. "github.com/onsi/gomega"    //nolint:revive
	"github.com/shopspring/decimal"
)

var _ = Context(
	"BeDecimalEquivalentTo", func() {
		Describe(
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
	},
)

var _ = Context(
	"BeDecimalGreaterThan", func() {
		Describe(
			"asserting on nil", func() {

				It(
					"actual is nil", func() {
						Ω(
							InterceptGomegaFailure(
								func() {
									Ω(nil).To(gdecimal.BeDecimalGreaterThan(decimal.Decimal{}))
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
									Ω(decimal.Decimal{}).To(gdecimal.BeDecimalGreaterThan(nil))
								},
							),
						).Error().To(Equal(errors.ErrNotBeNil))
					},
				)
			},
		)
	},
)

var _ = Context(
	"BeDecimalGreaterThanOrEqual", func() {
		Describe(
			"asserting on nil", func() {

				It(
					"actual is nil", func() {
						Ω(
							InterceptGomegaFailure(
								func() {
									Ω(nil).To(gdecimal.BeDecimalGreaterThanOrEqual(decimal.Decimal{}))
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
									Ω(decimal.Decimal{}).To(gdecimal.BeDecimalGreaterThanOrEqual(nil))
								},
							),
						).Error().To(Equal(errors.ErrNotBeNil))
					},
				)
			},
		)
	},
)

var _ = Context(
	"BeDecimalLessThan", func() {
		Describe(
			"asserting on nil", func() {

				It(
					"actual is nil", func() {
						Ω(
							InterceptGomegaFailure(
								func() {
									Ω(nil).To(gdecimal.BeDecimalLessThan(decimal.Decimal{}))
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
									Ω(decimal.Decimal{}).To(gdecimal.BeDecimalLessThan(nil))
								},
							),
						).Error().To(Equal(errors.ErrNotBeNil))
					},
				)
			},
		)
	},
)

var _ = Context(
	"BeDecimalLessThanOrEqual", func() {
		Describe(
			"asserting on nil", func() {

				It(
					"actual is nil", func() {
						Ω(
							InterceptGomegaFailure(
								func() {
									Ω(nil).To(gdecimal.BeDecimalLessThanOrEqual(decimal.Decimal{}))
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
									Ω(decimal.Decimal{}).To(gdecimal.BeDecimalLessThanOrEqual(nil))
								},
							),
						).Error().To(Equal(errors.ErrNotBeNil))
					},
				)
			},
		)
	},
)
