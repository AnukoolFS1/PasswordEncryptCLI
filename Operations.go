package main

import (
	"encoding/json"
	"errors"
	"os"
)

var Databasefile string = "passwords.json"

func Add(entry Entry) (string, error) {
	var entries []Entry

	data, err := os.ReadFile(Databasefile)

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

	// os.WriteFile(database, data, 0644)

	err := os.WriteFile(Databasefile, data, 0644)

	if err != nil {
		return "", err
	}
	return "File has been saved", nil //errors.New("Something wrong")
}

func Retrieve(username string) (Entry, error) {
	var userdata Entry
	data, err := os.ReadFile(Databasefile)
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

func List() ([]DuplicateTypeEntry, error) {
	data, err := os.ReadFile(Databasefile)
	if err != nil {
		return make([]DuplicateTypeEntry, 0), errors.New("An error occured during database reading")
	}

	var entries []DuplicateTypeEntry
	json.Unmarshal(data, &entries)

	return entries, nil
}

func UpdatePassword(username, password string) (string, error) {
	data, err := os.ReadFile(Databasefile)
	if err != nil {
		panic(err)
	}

	var NewEntries []Entry = make([]Entry, 0)
	var entries []Entry
	json.Unmarshal(data, &entries)

	for _, entry := range entries {
		if username == entry.Username {
			NewEntries = append(NewEntries,
				Entry{
					Username: username,
					Service:  entry.Service,
					Password: password})
		} else {
			NewEntries = append(NewEntries, entry)
		}
	}

	data, err = json.Marshal(NewEntries)
	if err != nil {
		return "", errors.New("Data has been corrupted")
	}

	err = os.WriteFile(Databasefile, data, 0644)

	return "Password has been changed", nil
}

func DeleteUser(username string) (string, error) {
	data, err := os.ReadFile(Databasefile)
	if err != nil {
		panic(err)
	}

	var NewEntries []Entry = make([]Entry, 0)
	var usernames = make(map[string]int)
	var entries []Entry
	json.Unmarshal(data, &entries)

	for _, entry := range entries {
		if username != entry.Username {
			NewEntries = append(NewEntries, entry)
		}
		usernames[entry.Username]++
	}
	if _, exist := usernames[username]; !exist {
		return "", errors.New("No Username has been found")
	}

	data, err = json.Marshal(NewEntries)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(Databasefile, data, 0644)

	return "User has been removed.", nil
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
