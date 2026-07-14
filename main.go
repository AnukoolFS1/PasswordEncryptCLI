package main

import (
	"encoding/json"
	// "errors"
	"fmt"
	"os"
)

type Entry struct {
	Service  string
	Username string
	Password string
}

func main() {
	// entries := basic()

	// fmt.Println()
	// result, err := saveData(entries)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(result)
	// }

	fmt.Println(RetrieveData())

}

func basic() []Entry {
	var entries []Entry

	entry := Entry{
		"Git", "Anukool", "123456",
	}

	entries = append(entries, entry)

	fmt.Println(entries)

	var service string
	var username string
	var password string

	fmt.Println("Servie: ")
	fmt.Scanln(&service)

	fmt.Println("Username: ")
	fmt.Scanln(&username)

	fmt.Println("Password: ")
	fmt.Scanln(&password)

	entry = Entry{Service: service, Username: username, Password: password}
	entries = append(entries, entry)

	return entries
}

func saveData(entries []Entry) (string, error) {
	data, _ := json.Marshal(entries)

	err := os.WriteFile("password.json", data, 0644)

	if err != nil {
		return "", err
	}
	return "File has been saved", nil //errors.New("Something wrong")
}

func RetrieveData() []Entry {
	data, err := os.ReadFile("password.json")

	var entries []Entry
	json.Unmarshal(data, &entries)

	if err != nil {
		panic(err)
	}
	return entries
}


