package calculator

import (
	"strconv"
	"strings"
)

func Calculate(input string) (string, error) {
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		panic("некорректный ввод. Используйте формат: 'a + b'")
	}

	a, b := tokens[0], tokens[2]
	operator := tokens[1]

	// Определяем, римские или арабские числа
	isRoman := isRoman(a) && isRoman(b)
	isArabic := isArabic(a) && isArabic(b)

	if !isRoman && !isArabic {
		panic("Используйте либо арабские, либо римские числа одновременно не < 1 и не > 10 включительно")
	}

	if isRoman {
		return calculateRoman(a, operator, b)
	}

	return calculateArabic(a, operator, b)
}

func calculateArabic(aStr, operator, bStr string) (string, error) {
	a, err := strconv.Atoi(aStr)
	if err != nil {
		panic("не удалось распознать число: " + aStr)
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		panic("не удалось распознать число: " + bStr)
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("числа должны быть от 1 до 10 включительно")
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		result = a / b
	default:
		panic("некорректный оператор: " + operator)
	}

	return strconv.Itoa(result), nil
}

func calculateRoman(aStr, operator, bStr string) (string, error) {
	a, err := RomanToArabic(aStr)
	if err != nil {
		return "", err
	}

	b, err := RomanToArabic(bStr)
	if err != nil {
		return "", err
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("римские числа должны быть от I до X включительно")
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		result = a / b
	default:
		panic("некорректный оператор: " + operator)
	}

	if result < 1 {
		panic("результат меньше единицы невозможен для римских чисел")
	}

	return ArabicToRoman(result), nil
}

// Функция для проверки, является ли строка римским числом
func isRoman(s string) bool {
	_, err := RomanToArabic(s)
	return err == nil
}

// Функция для проверки, является ли строка арабским числом
func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
