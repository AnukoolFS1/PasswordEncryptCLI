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
	test()
	testTwo()
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
