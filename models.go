package main

import (
	"fmt"
)

type EncryptedDatabase struct { // for data storage in encryption
	Salt []byte `json:"salt"`
	Data []byte `json:"data"`
}

type Entry struct {
	Service  string
	Username string
	Password string
}

func (entry Entry) String() string {

	return fmt.Sprintf("Username: %v\nService: %v\nPassword: %v\n", entry.Username, entry.Service, entry.Password)
}

func (entry Entry) SafeString() string {

	return fmt.Sprintf(
		"\nUsername: %s | Service: %s | Password: ",
		entry.Username,
		entry.Service)
}
