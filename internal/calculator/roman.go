package calculator

import "errors"

var romanToArabicMap = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
    "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRomanMap = []struct {
    Value  int
    Symbol string
}{
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func RomanToArabic(roman string) (int, error) {
    if value, exists := romanToArabicMap[roman]; exists {
        return value, nil
    }
    return 0, errors.New("некорректное римское число")
}

func ArabicToRoman(num int) string {
    var result string
    for _, entry := range arabicToRomanMap {
        for num >= entry.Value {
            result += entry.Symbol
            num -= entry.Value
        }
    }
    return result
}
