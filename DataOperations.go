package main

import (
	"encoding/json"
	"os"
)

func add(entry Entry) {
	var entries []Entry

	data, err := os.ReadFile("password.json")

	if err != nil {
		json.Unmarshal(data, &entries)

		defer SaveData(entries, entry)
	} else {
		SaveData(entries, entry)
	}
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

func Retrieve(username string) Entry {
	data, err := os.ReadFile("password.json")

	var entries []Entry
	json.Unmarshal(data, &entries)

	if err != nil {
		panic(err)
	}
	return entries
}

func List() []Entry {
	data, err := os.ReadFile("password.json")

	var entries []Entry
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
