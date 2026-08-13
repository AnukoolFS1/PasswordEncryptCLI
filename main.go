package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	// "os"
)

func main() {
	// TerminalArgs := os.Args

	// if !CheckArgs(TerminalArgs) || !ConfirmMaster() {
	// 	return
	// }

	// ArgsFunc(TerminalArgs)
	// test()
	// testTwo()
	// TestThree()
	// TestKey()
	TestFour()
}

func test() {
	entries := []Entry{
		{
			Service:  "GitHub",
			Username: "anukool",
			Password: "123456",
		},
	}

	data, _ := json.MarshalIndent(entries, "", "  ")

	fmt.Printf("%T\n", data)
	fmt.Println(data)
	fmt.Println(string(data))
}

func testTwo() {
	key := []byte("12345678901234567890123456789012")

	fmt.Println(key)
	fmt.Println(len(key))
	fmt.Printf("%T\n", key)
}

func TestKey() {
	// password := "IamOwner"

	salt := []byte("some-random-salt")

	// key := DeriveKey(password, salt)
	key1 := DeriveKey("IamOwner", salt)
	key2 := DeriveKey("WrongPassword", salt)

	fmt.Printf("Key: %x\n", key1)
	fmt.Println("Length:", len(key1))
	fmt.Printf("Key: %x\n", key2)
	fmt.Println("Length:", len(key2))
	same := string(key1) == string(key2)
	fmt.Println(same)
}

func TestFour() {
	salt := []byte("some-random-salt")
	key := DeriveKey("IamOwner", salt)

	entries := []Entry{
		{
			Service:  "GitHub",
			Username: "anukool",
			Password: "secret123",
		},
	}

	result, err := SaveEntries(entries, key, salt)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
