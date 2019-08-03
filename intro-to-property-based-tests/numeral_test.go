package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		given    int
	}{
		{"1 gets converted to I", "I", 1},
		{"2 gets converted to II", "II", 2},
		{"3 gets converted to III", "III", 3},
		{"4 gets converted to IV", "IV", 4},
		{"5 gets converted to V", "V", 5},
		{"6 gets converted to VI", "VI", 6},
		{"7 gets converted to VII", "VII", 7},
		{"8 gets converted to VIII", "VIII", 8},
		{"9 gets converted to IX", "IX", 9},
		{"10 gets converted to X", "X", 10},
		{"14 gets converted to XIV", "XIV", 14},
		{"18 gets converted to XVIII", "XVIII", 18},
		{"20 gets converted to XX", "XX", 20},
		{"39 gets converted to XXXIX", "XXXIX", 39},
		{"40 gets converted to XL", "XL", 40},
		{"47 gets converted to XLVII", "XLVII", 47},
		{"49 gets converted to XLIX", "XLIX", 49},
		{"50 gets converted to L", "L", 50},
		{"90 gets converted to XC", "XC", 90},
		{"100 gets converted to C", "C", 100},
		{"400 gets converted to CD", "CD", 400},
		{"500 gets converted to D", "D", 500},
		{"900 gets converted to CM", "CM", 900},
		{"1000 gets converted to M", "M", 1000},
		{"1984 gets converted to MCMLXXXIV", "MCMLXXXIV", 1984},
		{"3999 gets converted to MMMCMXCIX", "MMMCMXCIX", 3999},
		{"2019 gets converted to MMXIX", "MMXIX", 2019},
		{"1006 gets converted to MVI", "MVI", 1006},
		{"798 gets converted to DCCXCVIII", "DCCXCVIII", 798},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := ConvertToRoman(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %s, actual %s", tt.given, tt.expected, actual)
			}

		})
	}
}
