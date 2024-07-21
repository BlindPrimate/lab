package xor

import (
	"cryptopals/packages/hex_to_b64"
)

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
