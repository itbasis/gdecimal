package converter

import (
	"errors"

	gdecimalErrors "github.com/itbasis/gdecimal/errors"
	"github.com/shopspring/decimal"
)

func ToDecimal(value interface{}) (decimal.Decimal, error) {
	if value == nil {
		return decimal.Decimal{}, gdecimalErrors.ErrNotBeNil
	}

	if d, ok := value.(decimal.Decimal); ok {
		return d, nil
	}

	if s, ok := value.(string); ok {
		d, err := decimal.NewFromString(s) //nolint:varnamelen

		if err != nil {
			return d, errors.Join(err, gdecimalErrors.ErrFailedConvertStringToDecimal)
		}

		return d, nil
	}

	return decimal.Decimal{}, gdecimalErrors.ErrNotSupportedValue
}
