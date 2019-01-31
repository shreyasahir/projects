package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Encryption Program v 1.0")

	text := []byte("My super secret code stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("err")
	}

	// fmt.Println(gcm.Seal(nonce, nonce, text, nil))
	err = ioutil.WriteFile("mysecret.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		fmt.Println("Error writing to file", err)
	}
}
