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
		return make([]Entry, 0), errors.New("Database not found! Please create a new database first with command. PasswordEncrypt create Database")
	}

	err = json.Unmarshal(data, &entries)
	if err != nil {
		return make([]Entry, 0), errors.New("Data has been corrupted")
	}

	return entries, nil
}

func SaveEntries(entries []Entry) (string, error) {
	data, _ := json.Marshal(entries)

	// os.WriteFile(database, data, 0644)

	err := os.WriteFile(Databasefile, data, 0644)

	if err != nil {
		return "", err
	}
	return "File has been saved", nil //errors.New("Something wrong")
}

func CreateDb() (string, error) {
	_, err := os.Create(Databasefile)

	if err != nil {
		return "", err
	}
	return "File has been saved", nil
}
