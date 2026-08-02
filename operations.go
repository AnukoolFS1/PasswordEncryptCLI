package main

import (
	"errors"
)

func Add(entry Entry) (string, error) {
	entries, err := LoadEntries()

	if err != nil {
		return "", err
	}

	for _, users := range entries {
		if users.Username == entry.Username {
			return "", errors.New("UserName already exists \nPlease try with another username")
		}
	}

	entries = append(entries, entry)

	_, err = SaveEntries(entries)
	if err != nil {
		return "", err
	}

	return "Entry has been added.", nil
}

func Retrieve(username string) (Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return Entry{}, err
	}

	for _, entry := range entries {
		if entry.Username == username {
			return entry, nil
		}
	}
	return Entry{}, errors.New("User not found.")
}

func List() ([]Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return make([]Entry, 0), errors.New("An error occured during database reading")
	}

	return entries, nil
}

func UpdatePassword(username, password string) (string, error) {
	entries, err := LoadEntries()
	if err != nil {
		return "", err
	}

	var NewEntries []Entry = make([]Entry, 0)

	for _, entry := range entries {
		if username == entry.Username {
			entry.Password = password
			NewEntries = append(NewEntries, entry)
		} else {
			NewEntries = append(NewEntries, entry)
		}
	}

	_, erro := SaveEntries(NewEntries)
	if erro != nil {
		return "", erro
	}

	return "Password has been changed", err
}

func DeleteUser(username string) (string, error) {
	entries, err := LoadEntries()
	if err != nil {
		return "", err
	}

	var NewEntries []Entry = make([]Entry, 0)
	found := false

	for _, entry := range entries {
		if username == entry.Username {
			found = true
			continue
		}
		NewEntries = append(NewEntries, entry)
	}
	if !found {
		return "", errors.New("User not found.")
	}

	_, err = SaveEntries(NewEntries)
	if err != nil {
		return "", err
	}

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
