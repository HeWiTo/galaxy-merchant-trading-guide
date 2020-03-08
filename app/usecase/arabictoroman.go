package usecase

import (
	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/domain"
)

// ArabicToRoman exported
func ArabicToRoman(d int, arman []domain.ArabicRoman) string {
	var roman string

	for _, atr := range arman {
		for d >= atr.Arab {
			d -= atr.Arab
			roman += atr.Roman
		}
	}

	return roman
}
