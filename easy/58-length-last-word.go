package main

// 100% / 70%
func lengthOfLastWord(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			s = s[:i+1]
			break
		}
	}

	counter := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			counter += 1
		} else {
			break
		}
	}

	return counter
}
