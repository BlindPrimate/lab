package ciphers

import "cryptopals/packages/xor"

func RepeatingKeyCipher(text string, key string) []byte {
	textBytes := []byte(text)
	keyBytes := []byte(key)
	ciphertext := make([]byte, len(textBytes))
	windowLeft := 0
	windowRight := len(keyBytes) - 1

	for {
		if windowLeft == len(textBytes) {
			break
		}

		if windowLeft == windowRight && windowRight < len(textBytes) {
			windowRight += len(keyBytes)
		} else if windowLeft == windowRight {
			windowRight = len(textBytes)
		}

		keyIndex := windowLeft % len(keyBytes)
		ciphertext[windowLeft] = xor.SingleByteXor(textBytes[windowLeft], keyBytes[keyIndex])
		windowLeft++

	}

	return ciphertext
}
