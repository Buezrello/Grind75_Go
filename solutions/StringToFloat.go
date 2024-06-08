package solutions

import (
	"errors"
	"strings"
)

func StringToFloat(str string) (float64, error) {
	if len(strings.TrimSpace(str)) == 0 {
		return 0, errors.New("str is empty")
	}

	var result float64 = 0
	negative := false
	start := 0
	decimal := false
	var decimalDevisor float64 = 1

	if str[0] == '-' {
		negative = true
		start = 1
	}

	for i := start; i < len(str); i++ {
		if str[i] == '.' {
			if decimal {
				return 0, errors.New("second decimal point")
			}
			decimal = true
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			return 0, errors.New("Illegal character " + string(str[i]))
		}
		digit := float64(str[i] - '0')
		if decimal {
			decimalDevisor *= 10
			result += digit / decimalDevisor
		} else {
			result = result*10 + digit
		}
	}

	if negative {
		result = -result
	}

	return result, nil
}

//func getChar(str string, index int) rune {
//	return []rune(str)[index]
//}
