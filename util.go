package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func generateAppKey() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Panicln(err)
	}
	return fmt.Sprintf("base64:%s", base64.StdEncoding.EncodeToString(randomBytes))
}