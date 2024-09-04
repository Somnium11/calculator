package calculator

import (
	"strings"
)

var romanToArabicMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10, "L": 50, "C": 100,
}

var arabicToRomanMap = []struct {
	Value  int
	Symbol string
}{
	{100, "C"},
	{50, "L"},
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2, "II"},
	{1, "I"},
}

func RomanToArabic(roman string) (int, error) {
	if value, exists := romanToArabicMap[roman]; exists {
		return value, nil
	}
	panic("некорректный ввод")
}

func ArabicToRoman(num int) string {
	var result strings.Builder
	for _, entry := range arabicToRomanMap {
		for num >= entry.Value {
			result.WriteString(entry.Symbol)
			num -= entry.Value
		}
	}
	return result.String()
}
