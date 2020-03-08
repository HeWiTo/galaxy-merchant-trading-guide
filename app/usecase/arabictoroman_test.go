package usecase

import (
	"reflect"
	"testing"

	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/domain"
)

func TestArabicToRoman(t *testing.T) {
	tests := map[string]struct {
		d     int
		arman []domain.ArabicRoman
		want  string
	}{
		"TestCase1": {
			d: 2006,
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
			want: "MMVI",
		},
		"TestCase2": {
			d: 1944,
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
			want: "MCMXLIV",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ArabicToRoman(tc.d, tc.arman)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
