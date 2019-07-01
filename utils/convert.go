package utils

import "strconv"

// BytesToString : convert bytes to string
func BytesToString(bs []uint8) string {
	b := make([]byte, len(bs))

	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

// RContToInt : convert an []uint8 to int
func RContToInt(bs []uint8) int {
	ret, err := strconv.Atoi(BytesToString(bs))

	if err != nil {
		return (-1)
	}
	return (ret)
}
