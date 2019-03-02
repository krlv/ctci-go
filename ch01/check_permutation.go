package ch01

// CheckPermutationBySort returns true if one string is a permutation of the other
// Sorts both strings and checks if they are equal
func CheckPermutationBySort(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if SortString(s1) != SortString(s2) {
		return false
	}

	return true
}

// CheckPermutationBySet returns true if one string is a permutation of the other
// Use set of runes to reduce complexity
func CheckPermutationBySet(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	set := make(map[rune]int)

	for _, c := range s1 {
		set[c]++
	}

	for _, c := range s2 {
		set[c]--

		if set[c] < 0 {
			return false
		}
	}

	return true
}
