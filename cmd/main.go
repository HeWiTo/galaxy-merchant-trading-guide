package main

import (
	"fmt"

	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/domain"
	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/usecase"
)

func main() {
	arman := []domain.ArabicRoman{
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
	}

	gaman := []domain.GalaxyRoman{
		{1, "glob"},
		{5, "prok"},
		{10, "pish"},
		{50, "tegj"},
		{17, "silver"},
		{195.5, "iron"},
		{14450, "gold"},
	}

	arabicRoman := usecase.ArabicToRoman(1944, arman)
	romanArabic, _ := usecase.RomanToArabic("MMMVIII", arman)
	universeArabic, _ := usecase.UniverseToArabic("pish tegj glob glob", gaman)

	fmt.Println(arabicRoman)
	fmt.Println(romanArabic)
	fmt.Println(universeArabic)
}
