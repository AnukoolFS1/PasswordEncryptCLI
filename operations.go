package main

import (
	"errors"
)

func Add(entry Entry, Mpassword string) (string, error) {

	entries, salt, err := LoadEntries(Mpassword)
	if err != nil {
		return "", err
	}

	for _, users := range entries {
		if users.Username == entry.Username {
			return "", errors.New("UserName already exists \nPlease try with another username")
		}
	}

	entries = append(entries, entry)

	_, err = SaveEntries(entries, DeriveKey(Mpassword, salt), salt)
	if err != nil {
		return "", err
	}

	return "Entry has been added.", nil
}

func Retrieve(username string, Mpassword string) (Entry, error) {
	entries, _, err := LoadEntries(Mpassword)
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

func List(Mpassword string) ([]Entry, error) {
	entries, _, err := LoadEntries(Mpassword)
	if err != nil {
		return make([]Entry, 0), errors.New("An error occured during database reading")
	}

	return entries, nil
}

func UpdatePassword(username, password, Mpassword string) (string, error) {
	entries, salt, err := LoadEntries(Mpassword)
	if err != nil {
		return "", err
	}

	var NewEntries []Entry = make([]Entry, 0)
	found := false

	for _, entry := range entries {
		if username == entry.Username {
			entry.Password = password
			NewEntries = append(NewEntries, entry)
			found = true
		} else {
			NewEntries = append(NewEntries, entry)
		}
	}

	if !found {
		return "", errors.New("User not found.")
	}

	_, erro := SaveEntries(NewEntries, DeriveKey(Mpassword, salt), salt)
	if erro != nil {
		return "", erro
	}

	return "Password has been changed", err
}

func DeleteUser(username string, Mpassword string) (string, error) {
	entries, salt, err := LoadEntries(Mpassword)
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

	_, err = SaveEntries(NewEntries, DeriveKey(Mpassword, salt), salt)
	if err != nil {
		return "", err
	}

	return "User has been removed.", nil
}
