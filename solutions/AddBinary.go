package solutions

import "strconv"

func AddBinary(a string, b string) string {
	carry := 0
	result := ""
	var aLast int
	var bLast int
	for len(a) != 0 || len(b) != 0 {
		a, aLast = removeLast(a)
		b, bLast = removeLast(b)
		temp := (aLast + bLast + carry) % 2
		carry = (aLast + bLast + carry) / 2
		result = strconv.Itoa(temp) + result
	}
	if carry != 0 || len(result) == 0 {
		result = strconv.Itoa(carry) + result
	}

	return result
}

func removeLast(str string) (string, int) {
	if len(str) == 0 {
		return str, 0
	}

	last, _ := strconv.Atoi(str[len(str)-1:])
	str = str[0 : len(str)-1]

	return str, last
}
