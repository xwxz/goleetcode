package algorithm

import "math"

func IntegerToRoman(num int) string {
	var i2rMap = make(map[int]string)
	i2rMap[1] = "I"
	i2rMap[2] = "II"
	i2rMap[3] = "III"
	i2rMap[4] = "IV"
	i2rMap[5] = "V"
	i2rMap[6] = "VI"
	i2rMap[7] = "VII"
	i2rMap[8] = "VIII"
	i2rMap[9] = "IX"
	i2rMap[10] = "X"
	i2rMap[20] = "XX"
	i2rMap[30] = "XXX"
	i2rMap[40] = "XL"
	i2rMap[50] = "L"
	i2rMap[60] = "LX"
	i2rMap[70] = "LXX"
	i2rMap[80] = "LXXX"
	i2rMap[90] = "XC"
	i2rMap[100] = "C"
	i2rMap[200] = "CC"
	i2rMap[300] = "CCC"
	i2rMap[400] = "CD"
	i2rMap[500] = "D"
	i2rMap[600] = "DC"
	i2rMap[700] = "DCC"
	i2rMap[800] = "DCCC"
	i2rMap[900] = "CM"
	i2rMap[1000] = "M"
	i2rMap[2000] = "MM"
	i2rMap[3000] = "MMM"

	var tail int
	var roman string
	i := 0.0
	for num > 0 {
		tail = (num % 10) * int(math.Pow(10, i))
		roman = i2rMap[tail] + roman
		num = num / 10
		i++
	}
	return roman
}
