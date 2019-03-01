package ch01

// IsUnique determine if a string has all unique chars
func IsUnique(s string) bool {
	set := make(map[rune]bool)

	for _, c := range s {
		if set[c] {
			return false
		}

		set[c] = true
	}

	return true
}
