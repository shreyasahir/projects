package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("My Decryption program")
	key := []byte("passphrasewhichneedstobe32bytes!")
	cipherText, err := ioutil.ReadFile("mysecret.data")
	if err != nil {
		panic(err)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()

	if len(cipherText) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext))
}
