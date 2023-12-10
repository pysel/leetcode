package main

import "strings"

// 100% / 73%
func intToRoman(num int) string {
	romanString := ""

	if num > 1000 { // x > 1000
		thousands := num / 1000
		romanString += strings.Repeat("M", thousands)
		num -= 1000 * thousands
	}
	// num < 1000

	if num > 500 { // 500 < x < 1000
		if num < 1000 && num > 900 { // 900 < x < 1000
			romanString += "CM"
			num -= 900
		} else { // 500 < x < 900
			romanString += "D"
			num -= 500
		}
	}
	// num < 500

	if num > 100 { // 100 < x < 500
		if num < 500 && num > 400 { // 400 < x < 500
			romanString += "CD"
			num -= 400
		} else { // 100 < x < 400
			hundreds := num / 100
			romanString += strings.Repeat("C", hundreds)
			num -= 100 * hundreds
		}
	}
	// num < 100

	if num > 50 { // 50 < x < 100
		if num < 100 && num > 90 { // 90 < x < 100
			romanString += "XC"
			num -= 90
		} else { // 50 < x < 90
			romanString += "L"
			num -= 50
		}
	}
	// num < 50

	if num > 10 { // 10 < x < 50
		if num < 50 && num > 40 { // 40 < x < 50
			romanString += "XL"
			num -= 40
		} else { // 10 < x < 40
			tens := num / 10
			romanString += strings.Repeat("X", tens)
			num -= 10 * tens
		}
	}
	// num < 10

	if num >= 1 {
		if num == 9 {
			romanString += "IX"
			num -= 9
		} else if num == 4 {
			romanString += "IV"
			num -= 4
		} else {
			if num < 10 && num > 5 { // 5 < x < 10
				romanString += "V"
				num -= 5

				romanString += strings.Repeat("I", num)
			} else { // 1 < x < 5
				romanString += strings.Repeat("I", num)
			}
		}
	}

	return romanString
}
