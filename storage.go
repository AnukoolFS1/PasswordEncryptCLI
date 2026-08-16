package main

import (
	"encoding/json"
	"errors"
	"os"
)

var Databasefile string = "passwords.json"

func LoadEntries(password string) ([]Entry, []byte, error) {

	var database EncryptedDatabase

	// 1. Read the encrypted file
	fileData, err := os.ReadFile(Databasefile)
	if err != nil {
		return nil, nil, errors.New("database not found")
	}

	// 2. Decode the outer JSON
	err = json.Unmarshal(fileData, &database)
	if err != nil {
		return nil, nil, errors.New("database has been corrupted")
	}

	// 3. Derive the AES key using password + stored salt
	key := DeriveKey(password, database.Salt)

	// 4. Decrypt the encrypted data
	data, err := Decrypt(database.Data, key)
	if err != nil {
		return nil, nil, errors.New("incorrect master password or corrupted database")
	}

	// 5. Turn decrypted JSON back into []Entry
	var entries []Entry

	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, nil, errors.New("decrypted data is corrupted")
	}

	return entries, database.Salt, nil
}

func SaveEntries(entries []Entry, key []byte, salt []byte) (string, error) {

	// Convert Entries to JSON
	data, err := json.Marshal(entries)
	if err != nil {
		return "", err
	}

	// 2. Encrypt the JSON
	encrypted, err := Encrypt(data, key)
	if err != nil {
		return "", err
	}

	// 3. Put salt + encrypted data together
	database := EncryptedDatabase{
		Salt: salt,
		Data: encrypted,
	}

	// 4. Convert that structure to JSON
	fileData, err := json.Marshal(database)
	if err != nil {
		return "", err
	}

	// 5. Write it to disk
	err = os.WriteFile(Databasefile, fileData, 0644)
	if err != nil {
		return "", err
	}

	return "File has been saved", nil
}

func CreateDb(password string) (string, error) {
	data, _ := json.Marshal(make([]Entry, 0))
	err := os.WriteFile(Databasefile, data, 0644)
	salt := []byte("some-random-salt")

	SaveEntries(make([]Entry, 0), DeriveKey(password, salt), salt)

	if err != nil {
		return "", err
	}
	return "File has been saved", nil
}
