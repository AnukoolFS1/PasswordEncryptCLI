package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
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

func Decrypt(encrypted []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	if len(encrypted) < nonceSize {
		return nil, errors.New("encrypted data is too short")
	}

	nonce := encrypted[:nonceSize]
	ciphertext := encrypted[nonceSize:]

	data, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}
