package cnpj

import (
	"fmt"
	"regexp"
)

func checksum(ds []int32, ref []int32) int32 {
	var s int32
	for i, n := range ref {
		s += n * ds[i]
	}
	r := s % 11
	if r < 2 {
		return 0
	}
	return 11 - r
}

// IsValid checks whether CNPJ number is valid or not
func IsValid(n string) bool {
	u := Unmask(n)
	if len(u) != 14 {
		return false
	}

	ds := make([]int32, 14)
	for i, v := range []rune(u) {
		ds[i] = int32(v) - 48 // ASCII value minus 48 so "0" becomes 0 and so on
	}

	r1 := []int32{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	r2 := []int32{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	return checksum(ds, r1) == ds[12] && checksum(ds, r2) == ds[13]
}

// Mask returns the CNPJ number formatted
func Mask(n string) string {
	u := Unmask(n)
	if len(u) != 14 {
		return n
	}
	return fmt.Sprintf("%s.%s.%s/%s-%s", u[:2], u[2:5], u[5:8], u[8:12], u[12:])
}

// Unmask removes any non-digit (numeric) from the CNPJ number
func Unmask(n string) string {
	return regexp.MustCompile(`[^0-9A-Z]`).ReplaceAllString(n, "")
}

// Base returns the first 8 digits of the CNPJ number
func Base(n string) string {
	if !IsValid(n) {
		return ""
	}

	return Unmask(n)[0:8]
}

// Order returns the 9th, 10th, 11th and 12th digits of the CNPJ number.
func Order(n string) string {
	if !IsValid(n) {
		return ""
	}

	return Unmask(n)[8:12]
}

// CheckDigit returns the last 2 digits of the CNPJ number.
func CheckDigit(n string) string {
	if !IsValid(n) {
		return ""
	}

	return Unmask(n)[12:14]
}
