package main

import (
	"errors"
	"fmt"
)

func HandleAdd(args []string, password string) {
	if len(args) < 5 {
		SuccessOrError("", errors.New("Please provide:\nadd [service] [username] [password]"))
		return
	}
	fmt.Println("Adding data, Please Wait...")
	result, err := Add(Entry{args[2], args[3], args[4]}, password)
	SuccessOrError(result, err)
}

func HandleGet(args []string, password string) {
	if len(args) < 2 {
		SuccessOrError("", errors.New("Please provide username:\nget [username]"))
		return
	}
	fmt.Println("Retrieving data, Please Wait...")
	userdata, err := Retrieve(args[2], password)
	SuccessOrError(userdata.String(), err)
}

func HandleUpdate(args []string, password string) {
	if len(args) < 4 {
		fmt.Println("Please provide a username to delete entry.")
	} else {
		result, err := UpdatePassword(args[2], args[3], password)
		SuccessOrError(result, err)
	}
}

func HandleDelete(args []string, password string) {
	if len(args) < 3 {
		fmt.Println("Please provide a username to delete entry.")
	} else {
		result, err := DeleteUser(args[2], password)
		SuccessOrError(result, err)
	}
}

func HandleList(args []string, password string) {
	fmt.Println("Listing Data... Please Wait.")
	entries, err := List(password)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.SafeString())
	}
}

func SuccessOrError(success string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(success)
}
