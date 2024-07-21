package hex_to_b64

import (
	b64 "encoding/base64"
	"encoding/hex"
)

func EncodeHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func DecodeHex(src []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func EncodeBase64(input []byte) []byte {
	encoded := make([]byte, b64.StdEncoding.EncodedLen(len(input)))
	b64.StdEncoding.Encode(encoded, input)
	return encoded
}
