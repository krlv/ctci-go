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

// IsUniqueVanilla determine if a string has all unique chars
func IsUniqueVanilla(s string) bool {
	lens := len(s)
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}
