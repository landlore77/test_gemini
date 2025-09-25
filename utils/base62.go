package utils

import (
	"time"
)

const (
	base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base        = 62
)

func Base62Encode(num uint64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	encoded := ""
	for num > 0 {
		r := num % base
		num /= base
		encoded = string(base62Chars[r]) + encoded
	}
	return encoded
}

func GenerateID() string {
	// Generate a unique ID based on the current time in nanoseconds
	id := Base62Encode(uint64(time.Now().UnixNano()))
	if len(id) > 8 {
		return id[:8]
	}
	return id
}
