package xor

import (
	"cryptopals/packages/hex_to_b64"
)

func findMostCommonChar(chars map[byte]int) byte {
	var mostCommon byte
	mostCommonCount := 0
	for char, count := range chars {
		if count > mostCommonCount {
			mostCommon = char
			mostCommonCount = count
		}
	}
	return mostCommon
}

func XorBytes(src1 []byte, src2 []byte) ([]byte, error) {
	decoded1, err := hex_to_b64.DecodeHex(src1)
	if err != nil {
		return nil, err
	}

	decoded2, err := hex_to_b64.DecodeHex(src2)
	if err != nil {
		return nil, err
	}

	xor := make([]byte, len(decoded1))

	for i := 0; i < len(xor); i++ {
		xor[i] = decoded1[i] ^ decoded2[i]
	}

	return xor, nil
}

func FindXorChar(src []byte) byte {
	chars := map[byte]int{}
	for _, char := range src {
		if chars[char] == 0 {
			chars[char] = 1
		} else {
			chars[char]++
		}
	}
	mostCommon := findMostCommonChar(chars)
	return mostCommon
}

func SingleByteXor(input byte, key byte) byte {
	return input ^ key
}

func SingleByteXorOfBytes(input []byte, key byte) []byte {
	result := make([]byte, len(input))
	for i, val := range input {
		result[i] = SingleByteXor(val, key)
	}
	return result
}
