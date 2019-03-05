package ch01

import "strconv"

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

// CompressBytes returns "compressed" string - all duplicates chars relpaced with "xN",
// where x is original char and N is a number of chars
// If compressed string isn't smaller, than original, original string will be returned
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
