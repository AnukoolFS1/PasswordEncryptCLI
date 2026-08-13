package main

import (
	"encoding/json"
	"errors"
	"os"
)

var Databasefile string = "passwords.json"

func LoadEntries() ([]Entry, error) {
	var entries []Entry

	data, err := os.ReadFile(Databasefile)
	if err != nil {
		return make([]Entry, 0), errors.New("Database not found! Please create a new database first with command. PasswordEncrypt create-db")
	}

	err = json.Unmarshal(data, &entries)
	if err != nil {
		return make([]Entry, 0), errors.New("Data has been corrupted")
	}

	return entries, nil
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

func CreateDb() (string, error) {
	data, _ := json.Marshal(make([]Entry, 0))
	err := os.WriteFile(Databasefile, data, 0644)

	if err != nil {
		return "", err
	}
	return "File has been saved", nil
}
