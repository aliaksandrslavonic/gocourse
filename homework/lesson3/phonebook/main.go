package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type PhoneBook struct {
	Name     string
	LastName string
	Phone    string
}

func main() {
	phone := PhoneBook{"Alex", "Bylina", "+3752900000000"}

	payload, err := json.Marshal(phone)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(payload))
	ioutil.WriteFile("phonelist.json", payload, 0644)

}
