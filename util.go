package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateAppKey() (result string, err error) {
	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return
	}
	result = fmt.Sprintf("base64:%s", base64.StdEncoding.EncodeToString(randomBytes))
	return
}