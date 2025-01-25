package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int
	Balance int
}

type AccountZ struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))

	// Encode json to stdout
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	jsonPure := []byte(`{"Number":2,"Balance":100}`)
	var accountX Account

	// Decode json to struct
	err = json.Unmarshal(jsonPure, &accountX)
	if err != nil {
		panic(err)
	}
	println(accountX.Number)

	jsonPureIncompatible := []byte(`{"n":3,"b":100}`)
	var accountZ AccountZ

	err = json.Unmarshal(jsonPureIncompatible, &accountZ)
	if err != nil {
		panic(err)
	}
	println(accountZ.Number)
}
