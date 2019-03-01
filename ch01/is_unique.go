package ch01

// IsUnique determine if a string has all unique chars
// This solution is based on HashTable (Go map) data structure
func IsUnique(s string) bool {
	set := make(map[rune]bool)

	for _, c := range s {
		if _, ok := set[c]; ok {
			return false
		}

		set[c] = true
	}

	return true
}

// IsUniqueVanilla determine if a string has all unique chars
// This solution doesn't utilize any additional data structures
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
