package usecase

import (
	"reflect"
	"testing"

	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/domain"
)

func TestRomanToArabic(t *testing.T) {
	tests := map[string]struct {
		v     string
		arman []domain.ArabicRoman
		want  int
	}{
		"TestCase1": {
			v: "MMVI",
			arman: []domain.ArabicRoman{
				{1000, "M"},
				{900, "CM"},
				{500, "D"},
				{400, "CD"},
				{100, "C"},
				{90, "XC"},
				{50, "L"},
				{40, "XL"},
				{10, "X"},
				{5, "V"},
				{4, "IV"},
				{1, "I"},
			},
			want: 2006,
		},
		"TestCase2": {
			v: "MCMXLIV",
			arman: []domain.ArabicRoman{
				{1000, "M"},
				{900, "CM"},
				{500, "D"},
				{400, "CD"},
				{100, "C"},
				{90, "XC"},
				{50, "L"},
				{40, "XL"},
				{10, "X"},
				{5, "V"},
				{4, "IV"},
				{1, "I"},
			},
			want: 1944,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := RomanToArabic(tc.v, tc.arman)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
