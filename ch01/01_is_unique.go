package ch01

import "strings"

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

// IsUniqueBits determines if a string has all unique chars (only works for ASCII)
// This solution doesn't utilize any additional data structures and has O(n) time complexity
func IsUniqueBits(s string) bool {

	s = strings.ToLower(s)
	var vector int32
	for _, rune := range s {
		index := rune - 'a'
		mask := int32(1 << index)
		if (vector & mask) == mask {
			return false
		}
		vector = vector | mask
	}
	return true
}
