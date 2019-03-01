package ch01

import "sort"

// SortableString is a representation of a sortable string
type SortableString []rune

func (s SortableString) Len() int           { return len(s) }
func (s SortableString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortableString) Less(i, j int) bool { return s[i] < s[j] }

// SortString returns the sorted string
func SortString(s string) string {
	sr := SortableString(s)
	sort.Sort(sr)

	return string(sr)
}
