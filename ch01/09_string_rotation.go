package ch01

import "strings"

// IsRotation returns true if the first string is a rotation of the scond one
func IsRotation(s1 string, s2 string) bool {
	s1len, s2len := len(s1), len(s2)
	if s1len != s2len {
		return false
	}

	s2 += s2
	return strings.Contains(s2, s1)
}
