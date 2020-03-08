package usecase

import (
	"regexp"

	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/domain"
)

// UniverseToArabic exported
func UniverseToArabic(v string, gaman []domain.GalaxyRoman) (float32, error) {
	var (
		arabic, nominal, temp float32
		word                  string
	)

	lenString := len(v)
	words := regexp.MustCompile(" ").Split(v, lenString)

	for _, gar := range gaman {
		for _, word = range words {
			if word == gar.Roman {
				if gar.Roman == "gold" {
					nominal = gar.Universe
				} else if gar.Roman == "iron" {
					nominal = gar.Universe
				} else if gar.Roman == "silver" {
					nominal = gar.Universe
				} else {
					if temp > gar.Universe || temp == gar.Universe {
						arabic += gar.Universe
					} else if temp < gar.Universe {
						arabic = gar.Universe - arabic
					}
					temp = gar.Universe
				}
			}
		}

		if nominal == 0 {
			nominal++
		}
	}
	arabic *= nominal

	return arabic, nil
}
