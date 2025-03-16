package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)


func main() {
	const apiKey = "1234567890"
	const apiSecret = "0987654321"

	h := hmac.New(sha256.New, []byte(apiSecret))
	fmt.Println(h)
}
