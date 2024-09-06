package calculator_test

import (
	"calculator/internal/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)


// Корректность вычислений с римскими числами
func TestCalculateRoman(t *testing.T) {
	assert := assert.New(t)

	result, err := calculator.Calculate("IV + VI")
	assert.NoError(err)
	assert.Equal("X", result)

	result, err = calculator.Calculate("VII - III")
	assert.NoError(err)
	assert.Equal("IV", result)

	result, err = calculator.Calculate("II * III")
	assert.NoError(err)
	assert.Equal("VI", result)

	result, err = calculator.Calculate("X / II")
	assert.NoError(err)
	assert.Equal("V", result)
}


// Корректность вычислений с арабскими числами
func TestCalculateArabic(t *testing.T) {
	assert := assert.New(t)

	result, err := calculator.Calculate("5 + 3")
	assert.NoError(err)
	assert.Equal("8", result)

	result, err = calculator.Calculate("7 - 2")
	assert.NoError(err)
	assert.Equal("5", result)

	result, err = calculator.Calculate("4 * 3")
	assert.NoError(err)
	assert.Equal("12", result)

	result, err = calculator.Calculate("9 / 3")
	assert.NoError(err)
	assert.Equal("3", result)
}

// Проверка на выход за пределы диапазона
func TestCalculateOutOfRange(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		calculator.Calculate("11 + 5")
	})
	
	assert.Panics(func() {
		calculator.Calculate("5 + 12")
	})
}

// Проверка на смешивание типов
func TestCalculateMixedNumerals(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		calculator.Calculate("5 + V")
	})

	assert.Panics(func() {
		calculator.Calculate("IV * 3")
	})
}

// Проверка деления на ноль
func TestCalculateDivisionByZero(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		calculator.Calculate("5 / 0")
	})

	assert.Panics(func() {
		calculator.Calculate("X / 0")
	})
}

// проверка на некорректные входные данные
func TestCalculateInvalidInput(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		calculator.Calculate("5 +")
	})

	assert.Panics(func() {
		calculator.Calculate("5")
	})

	assert.Panics(func() {
		calculator.Calculate("5 + 5 + 5")
	})
}


// Проверка на выход за пределы диапазона
func TestCalculateRomanNegativeResult(t *testing.T) {
	assert := assert.New(t)

	// Тестирование случаев, когда результат меньше единицы
	assert.Panics(func() {
		calculator.Calculate("IV - V")
	})

	assert.Panics(func() {
		calculator.Calculate("I - II")
	})

	assert.Panics(func() {
		calculator.Calculate("V / VI")
	})
}