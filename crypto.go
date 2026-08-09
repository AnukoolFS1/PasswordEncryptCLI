package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func Encrypt(data []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	cyphertext := gcm.Seal(nil, nonce, data, nil)

	encrypted := append(nonce, cyphertext...)

	return encrypted, nil

	// return cyphertext, nil
}


func TestThree () {
	key := []byte("12345678901234567890123456789012")

	message := []byte("Hello, Anukool!")

	encrypted, err := Encrypt(message, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("Original:", string(message))
	fmt.Println("Encrypted:", encrypted)
}