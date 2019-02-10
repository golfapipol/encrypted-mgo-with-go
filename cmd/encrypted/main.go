package main

import (
	"encryptedmgo/util"
	"log"
)

const (
	key = "6368616e676520746869732070617373"
)

func main() {
	input := "exampleplaintext"
	log.Println("input is:", input)
	encryptedText, err := util.Encrypt(key, input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("encryptedText: %s\n", encryptedText)

	decryptedText, err := util.Decrypt(key, encryptedText)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("decryptedText: %s\n", decryptedText)
}
