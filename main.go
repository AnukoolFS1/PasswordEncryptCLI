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