package main

import (
	"encoding/json"
	"errors"
	"os"
)

type DuplicateTypeEntry Entry

func add(entry Entry) (string, error) {
	var entries []Entry

	data, err := os.ReadFile("password.json")

	if err != nil {
		SaveData(entries, entry)
	} else {
		err := json.Unmarshal(data, &entries)
		if err != nil {
			return "", errors.New("Data has been corrupted.")
		}

		for _, users := range entries {
			if users.Username == entry.Username {
				return "", errors.New("UserName already exists \nPlease try with another username")
			}
		}

		SaveData(entries, entry)
	}
	return "Entry has been added.", nil
}

func SaveData(entries []Entry, entry Entry) (string, error) {
	entries = append(entries, entry)

	data, _ := json.Marshal(entries)

	// os.WriteFile("password.json", data, 0644)

	err := os.WriteFile("password.json", data, 0644)

	if err != nil {
		return "", err
	}
	return "File has been saved", nil //errors.New("Something wrong")
}

func Retrieve(username string) (Entry, error) {
	var userdata Entry
	data, err := os.ReadFile("password.json")
	if err != nil {
		return Entry{}, err
	}
	// if err != nil {
	// 	panic(err)
	// }

	var entries []Entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return Entry{}, errors.New("Data has been corrupted.")
	}

	for _, entry := range entries {
		if entry.Username == username {
			userdata = entry
		}
	}
	return userdata, nil
}

func List() []DuplicateTypeEntry {
	data, err := os.ReadFile("password.json")

	var entries []DuplicateTypeEntry
	json.Unmarshal(data, &entries)

	if err != nil {
		panic(err)
	}
	return entries
}

// func basic() []Entry {
// 	var entries []Entry

// 	entry := Entry{
// 		"Git", "Anukool", "123456",
// 	}

// 	entries = append(entries, entry)

// 	fmt.Println(entries)

// 	var service string
// 	var username string
// 	var password string

// 	fmt.Println("Servie: ")
// 	fmt.Scanln(&service)

// 	fmt.Println("Username: ")
// 	fmt.Scanln(&username)

// 	fmt.Println("Password: ")
// 	fmt.Scanln(&password)

// 	entry = Entry{Service: service, Username: username, Password: password}
// 	entries = append(entries, entry)

// 	return entries
// }
