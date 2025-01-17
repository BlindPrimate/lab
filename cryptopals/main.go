package main

import (
	"cryptopals/packages/ciphers"
	"cryptopals/packages/hex_to_b64"
	"fmt"
)

func main() {
	// src := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	// bits, err := hex_to_b64.DecodeHex(src)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// mostCommon := xor.FindXorChar(src)
	// fmt.Println(mostCommon)

	// for i := range 256 {
	// 	trans := xor.SingleByteXorOfBytes(bits, byte(i))
	// 	fmt.Println(string(trans))
	// 	fmt.Println(i)
	// }
	text := `Burning 'em, if you ain't quick and nimble
			I go crazy when I hear a cymbal`

	ciphertext := ciphers.RepeatingKeyCipher(text, "ICE")

	fmt.Println(string(hex_to_b64.EncodeHex(ciphertext)))

}
