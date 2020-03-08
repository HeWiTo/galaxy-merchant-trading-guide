package usecase

import (
	"errors"
	"strings"

	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/domain"
)

// RomanToArabic exported
func RomanToArabic(v string, arman []domain.ArabicRoman) (int, error) {
	var arabic, m, c, x, i int

	for _, r := range v {
		if r == 'M' {
			m++
			if m > 3 {
				return arabic, errors.New("M can not be repeated more than three times")
			}
		}

		if r == 'C' {
			c++
			if c > 3 {
				return arabic, errors.New("C can not be repeated more than three times")
			}
		}

		if r == 'X' {
			x++
			if x > 3 {
				return arabic, errors.New("X can not be repeated more than three times")
			}
		}

		if r == 'I' {
			i++
			if i > 3 {
				return arabic, errors.New("I can not be repeated more than three times")
			}
		}
	}

	for _, rta := range arman {
		for strings.HasPrefix(v, rta.Roman) {
			v = v[len(rta.Roman):]
			arabic += rta.Arab
		}
	}

	return arabic, nil
}
