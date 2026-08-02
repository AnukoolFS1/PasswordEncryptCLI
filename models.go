package main

import (
	"fmt"
)

type Entry struct {
	Service  string
	Username string
	Password string
}

type DuplicateTypeEntry Entry


func (entry Entry) String() string {

	return fmt.Sprintf("Username: %v\nService: %v\nPassword: %v\n", entry.Username, entry.Service, entry.Password)
}

func (entry DuplicateTypeEntry) String() string {

	return fmt.Sprintf(
		"\nUsername: %s | Service: %s | Password: %s",
		entry.Username,
		entry.Service,
		entry.Password)
}
