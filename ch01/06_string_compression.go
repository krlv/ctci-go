package ch01

import (
	"strconv"
	"strings"
)

// Compress returns "compressed" string - all duplicates chars relpaced with "xN",
// where x is original char and N is a number of chars
// If compressed string isn't smaller, than original, original string will be returned
func Compress(s string) string {
	slen := len(s)

	// if len is less than 3, the original string will be smaller
	if slen < 3 {
		return s
	}

	compressed := ""
	count := 1
	for i := 1; i < slen; i++ {
		if s[i-1] != s[i] {
			compressed += string(s[i-1]) + strconv.Itoa(count)
			count = 1
		} else {
			count++
		}
	}
	compressed += string(s[slen-1]) + strconv.Itoa(count)

	if len(compressed) < slen {
		return compressed
	}

	return s
}

// CompressBytes is the Compress implementation using byte slice
func CompressBytes(s string) string {
	slen := len(s)

	// if len is less than 3, the original string will be smaller
	if slen < 3 {
		return s
	}

	bytes := make([]byte, 0)
	count := 1
	for i := 1; i < slen; i++ {
		if s[i-1] != s[i] {
			bytes = append(bytes, s[i-1])
			bytes = append(bytes, []byte(strconv.Itoa(count))...)
			count = 1
		} else {
			count++
		}
	}
	bytes = append(bytes, s[slen-1])
	bytes = append(bytes, []byte(strconv.Itoa(count))...)

	if len(bytes) < slen {
		return string(bytes)
	}

	return s
}

// CompressBuilder is the Compress implementation based on strings.Builder
func CompressBuilder(s string) string {
	slen := len(s)

	// if len is less than 3, the original string will be smaller
	if slen < 3 {
		return s
	}

	var compressed strings.Builder
	count := 1
	for i := 1; i < slen; i++ {
		if s[i-1] != s[i] {
			compressed.WriteByte(s[i-1])
			compressed.WriteString(strconv.Itoa(count))

			count = 1
		} else {
			count++
		}
	}
	compressed.WriteByte(s[slen-1])
	compressed.WriteString(strconv.Itoa(count))

	if compressed.Len() < slen {
		return compressed.String()
	}

	return s
}

// CompressBuilderCap is the Compress implementation based on strings.Builder
// To optimize strings.Builder capacity, we're going to calculate the length
// pf compressed string ahead
func CompressBuilderCap(s string) string {
	slen := len(s)
	cmplen := compressedLength(s)

	// if len is less than compressed length, the original string will be returned
	if slen <= cmplen {
		return s
	}

	var compressed strings.Builder
	compressed.Grow(cmplen)

	count := 1
	for i := 1; i < slen; i++ {
		if s[i-1] != s[i] {
			compressed.WriteByte(s[i-1])
			compressed.WriteString(strconv.Itoa(count))

			count = 1
		} else {
			count++
		}
	}
	compressed.WriteByte(s[slen-1])
	compressed.WriteString(strconv.Itoa(count))

	return compressed.String()
}

// compressedLength returns compressed length of string s
func compressedLength(s string) int {
	slen := len(s)
	cmlen := 0
	count := 1
	for i := 1; i < slen; i++ {
		if s[i-1] != s[i] {
			cmlen += 1 + len(strconv.Itoa(count))
			count = 1
		} else {
			count++
		}
	}
	return cmlen + 1 + len(strconv.Itoa(count))
}
