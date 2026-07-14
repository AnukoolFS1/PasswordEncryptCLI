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
	// result, err := SaveData(entries)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(result)
	// }

	fmt.Println(RetrieveData())

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


