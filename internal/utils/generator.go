package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

const _base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Approach 1: Hashing and Encoding

// A common approach for generating short URLs is to use a hash function, such as MD5 or SHA-256 to generate
//
//	a fixed-length hash of the original URL.
//
// This hash is then encoded into a shorter form using Base62.

func GenerateShortURL(longURL string) string {
	// Step 1: generate MD5 hash of the longURL
	hash := md5.Sum([]byte(longURL))

	// Step 2: convert first 6 bytes of the hash to hexadecimal string
	first6BytesHex := hex.EncodeToString(hash[:6])

	// Step 3: convert the hexadecimal string to a decimal number
	decimalValue := new(big.Int)
	decimalValue.SetString(first6BytesHex, 16)

	// Step 4: encode the decimal number to a Base62 string
	shortURL := base62Encode(decimalValue)

	return shortURL
}

func base62Encode(num *big.Int) string {
	var result string
	base := big.NewInt(62)

	for num.Cmp(big.NewInt(0)) > 0 {
		remainder := new(big.Int)
		num.DivMod(num, base, remainder)
		result = string(_base62Chars[remainder.Int64()]) + result
	}

	return result
}
