package main

import (
	"cryptopals/packages/hex_to_b64"
	"cryptopals/packages/xor"
	"fmt"
)

func main() {
	src1 := []byte("1c0111001f010100061a024b53535009181c")
	src2 := []byte("686974207468652062756c6c277320657965")
	xorResult, err := xor.XorBytes(src1, src2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(hex_to_b64.EncodeHex(xorResult)))
}
