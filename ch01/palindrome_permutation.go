package ch01

// IsPalindromePermutation returns true if a srting is a palindrome permutation
func IsPalindromePermutation(s string) bool {
	set := make(map[rune]int)

	for _, c := range s {
		if c == ' ' {
			continue
		}

		set[c]++
	}

	odd := false
	for _, n := range set {
		if n%2 != 0 {
			if odd {
				return false
			}

			odd = true
		}
	}

	return true
}
