package main

import (
	"calculator/internal/calculator"
	"fmt"
)

func main() {
	c := &calculator.Calculator{}
	var num1, num2 int
	fmt.Scan("Введите 2 числа через пробел: ", &num1, &num2)
	op := "+"

	if op == "+" {
        fmt.Println(c.Addition(num1, num2))
    } else if op == "-" {
        fmt.Println(c.Subtraction(num1, num2))
    } else if op == "*" {
        fmt.Println(c.Multiplication(num1, num2))
    } else if op == "/" {
        fmt.Println(c.Division(num1, num2))
    } else {
        fmt.Println("Invalid operator")
    }
}
