package calculator

import "fmt"

type Calculator struct{}

func (c *Calculator) Addition(num1, num2 int) int {
	return num1 + num2
}
// func (c *Calculator) StringAddition(str1, str2 string) string {
// 	return str1 + str2
	
// }
func (c *Calculator) Subtraction(num1, num2 int) int {
	return num1 - num2
}

// func (c *Calculator) StringSubtraction(str1, str2 string) string {
// 	return str1 + str2
// }

func (c *Calculator) Multiplication(num1, num2 int) int {
	return num1 * num2
}

// func (c *Calculator) StringMultiplication(str1, str2 string) string {
// 	return str1 + str2
// }

func (c *Calculator) Division(num1, num2 int) int {
	if num2 == 0 {
		fmt.Println("На 0 делить нельзя")
	}
	return num1 / num2
}

// func (c *Calculator) StringDivision(str1, str2 string) string {
// 	return str1 + str2
// }	