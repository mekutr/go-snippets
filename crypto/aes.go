package main

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	mrand "math/rand"
	"net/url"
	"time"
)

const cAlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const myMessage = "my very long sample message"

func main() {
	iv, err := createInitVector()
	if err != nil {
		panic(err)
	}

	key := []byte(randString(32))
	blockCipher, err := createBlockCipher(key)
	if err != nil {
		panic(err)
	}

	// It's important to remember that cipherTexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	cipherText := encryptCTRAES(myMessage, iv, blockCipher)
	fmt.Println(cipherText)

	b64e := base64Encoded(cipherText)
	fmt.Println(b64e)

	qe := queryEscaped(b64e)
	fmt.Println(qe)

	qu := queryUnescaped(qe)
	fmt.Println(qu)

	b64d := base64Decoded(qu)
	fmt.Println(b64d)

	plainText := decryptCTRAES(b64d, blockCipher)
	fmt.Println(plainText)

	fmt.Printf("%s\n", plainText)
}

func queryEscaped(s string) string {
	return url.QueryEscape(s)
}

func queryUnescaped(s string) string {
	unescaped, _ := url.QueryUnescape(s)
	return unescaped
}

func base64Encoded(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func base64Decoded(s string) []byte {
	decoded, _ := base64.StdEncoding.DecodeString(s)
	return decoded
}

func createBlockCipher(key []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return block, err
}

func createInitVector() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(crand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}

func encryptCTRAES(textToEncrypt string, iv []byte, block cipher.Block) []byte {
	byteForm := []byte(textToEncrypt)
	var cipherText []byte
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	cipherText = append(cipherText, iv...)
	cipherText = append(cipherText, byteForm...)
	stream := cipher.NewCTR(block, iv)
	// It's important to remember that cipherTexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteForm)
	return cipherText
}

func decryptCTRAES(textToDecrypt []byte, block cipher.Block) []byte {
	byteForm := textToDecrypt
	plainText := make([]byte, len(byteForm)-aes.BlockSize)
	iv := byteForm[:aes.BlockSize]
	// CTR mode is the same for both encryption and decryption, so we can
	// also decrypt that cipherText with NewCTR.
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plainText, byteForm[aes.BlockSize:])
	return plainText
}

func randString(n int) string {
	mrand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = cAlphaNumeric[mrand.Intn(len(cAlphaNumeric))]
	}
	return string(b)
}
