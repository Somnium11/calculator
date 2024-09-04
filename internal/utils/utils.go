package utils

func intToRoman(num int) string {
	val := []int{1,2,3,4,5,6,7,8,9,10}
	syb := []string{"I","II","III","IV","V","VI","VII","VIII","IX","X"}

	var roman string
	for i := 0; num > 0; i++ {
		for num >= val[i] {
			roman += syb[i]
			num -= val[i]
		}
	}

	return roman
}
