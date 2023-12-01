package main

// R: 67.5%; M: 88.7%
func romanToInt(s string) int {
	mapping := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if len(s) == 1 {
		return mapping[s]
	}

	output := 0
	for i := 0; i < len(s)-1; i++ {
		if mapping[string(s[i])] >= mapping[string(s[i+1])] { // a normal value
			output += mapping[string(s[i])]
		} else {
			output += mapping[string(s[i+1])] - mapping[string(s[i])]
			i++
		}
	}

	// cover last value
	if mapping[string(s[len(s)-1])] <= mapping[string(s[len(s)-2])] {
		output += mapping[string(s[len(s)-1])]
	}

	return output
}
