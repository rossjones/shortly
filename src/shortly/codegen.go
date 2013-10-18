package main

import "bytes"
import "strings"

var encoding_chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func base62_encode(number uint64) string {
	result := bytes.NewBufferString("         ")

	for number >= 62 {
		r := number % 62
		result.WriteString(string(encoding_chars[r]))
		number = number / 62
	}

	result.WriteString(string(encoding_chars[number]))
	return reverse(strings.TrimSpace(result.String()))
}

func base62_decode(hash string) (final uint64) {

	for i := 0; i < len(hash); i++ {
		value := uint64(strings.Index(encoding_chars, string(hash[i])))
		num := value * kapow(uint64(62), uint64(len(hash)-(i+1)))
		final += num
	}

	return
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func kapow(x, y uint64) (r uint64) {
	if x == r {
		return
	}
	r = 1
	if x == r {
		return
	}
	for y > 0 {
		if y&1 == 1 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return
}
