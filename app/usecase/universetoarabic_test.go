package usecase

import (
	"reflect"
	"testing"

	"github.com/HeWiTo/galaxy-merchant-trading-guide/app/domain"
)

func TestUniverseToArabic(t *testing.T) {
	tests := map[string]struct {
		v     string
		gaman []domain.GalaxyRoman
		want  float32
	}{
		"TestCase1": {
			v: "pish tegj glob glob",
			gaman: []domain.GalaxyRoman{
				{1, "glob"},
				{5, "prok"},
				{10, "pish"},
				{50, "tegj"},
				{17, "silver"},
				{195.5, "iron"},
				{14450, "gold"},
			},
			want: 42,
		},
		"TestCase2": {
			v: "glob glob silver",
			gaman: []domain.GalaxyRoman{
				{1, "glob"},
				{5, "prok"},
				{10, "pish"},
				{50, "tegj"},
				{17, "silver"},
				{195.5, "iron"},
				{14450, "gold"},
			},
			want: 34,
		},
		"TestCase3": {
			v: "glob prok silver",
			gaman: []domain.GalaxyRoman{
				{1, "glob"},
				{5, "prok"},
				{10, "pish"},
				{50, "tegj"},
				{17, "silver"},
				{195.5, "iron"},
				{14450, "gold"},
			},
			want: 68,
		},
		"TestCase4": {
			v: "glob prok gold",
			gaman: []domain.GalaxyRoman{
				{1, "glob"},
				{5, "prok"},
				{10, "pish"},
				{50, "tegj"},
				{17, "silver"},
				{195.5, "iron"},
				{14450, "gold"},
			},
			want: 57800,
		},
		"TestCase5": {
			v: "pish pish iron",
			gaman: []domain.GalaxyRoman{
				{1, "glob"},
				{5, "prok"},
				{10, "pish"},
				{50, "tegj"},
				{17, "silver"},
				{195.5, "iron"},
				{14450, "gold"},
			},
			want: 3910,
		},
		"TestCase6": {
			v: "glob prok iron",
			gaman: []domain.GalaxyRoman{
				{1, "glob"},
				{5, "prok"},
				{10, "pish"},
				{50, "tegj"},
				{17, "silver"},
				{195.5, "iron"},
				{14450, "gold"},
			},
			want: 782,
		},
		"TestCase7": {
			v: "wood could a woodchuck chuck if a woodchuck could chuck wood",
			gaman: []domain.GalaxyRoman{
				{1, "glob"},
				{5, "prok"},
				{10, "pish"},
				{50, "tegj"},
				{17, "silver"},
				{195.5, "iron"},
				{14450, "gold"},
			},
			want: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := UniverseToArabic(tc.v, tc.gaman)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
