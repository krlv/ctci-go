package ch01

// CheckPermutation returns true if one string is a permutation of the other
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if SortString(s1) != SortString(s2) {
		return false
	}

	return true
}
