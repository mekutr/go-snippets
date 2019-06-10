package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"math/rand"
	"net/url"
	"time"
)

const cAlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const myMessage = "my very long sample message"

func main() {
	key := randString(2048)

	mac := newHmacForSha2(key)
	mac.Write([]byte(myMessage))

	hash := mac.Sum(nil)
	fmt.Println(hash)

	b64e := base64Encoded(hash)
	fmt.Println(b64e)

	qe := queryEscaped(b64e)
	fmt.Println(qe)

	qu := queryUnescaped(qe)
	fmt.Println(qu)

	b64d := base64Decoded(qu)
	fmt.Println(b64d)

	macsEqual := hmacEqual(hash, b64d)
	fmt.Println(macsEqual)
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

// this function creates the hmac with sha2(512)
func newHmacForSha2(key string) hash.Hash {
	return hmac.New(sha512.New, []byte(key))
}

// hmacEqual identifies whether b1 and b2 match
func hmacEqual(b1 []byte, b2 []byte) bool {
	return hmac.Equal(b1, b2)
}

// This creates the authentication key to create the MAC
func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = cAlphaNumeric[rand.Intn(len(cAlphaNumeric))]
	}
	return string(b)
}
