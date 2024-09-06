package main

import (
	"bufio"
	"calculator/internal/calculator"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа и операцию (например 3 + 5 или IV * VI). Для выхода введите 'exit':")

	for {
		fmt.Print("> ") // Приглашение к вводу
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		result, err := calculator.Calculate(input)
		if err != nil {
			fmt.Println("Ошибка:", err) // Вывод ошибки
		} else {
			fmt.Println("Результат:", result) // Вывод результата
		}
	}
}
