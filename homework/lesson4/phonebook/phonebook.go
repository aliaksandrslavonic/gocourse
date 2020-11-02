package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	name   string
	number string
}

type Contacts []*Contact

func main() {

	contacts := []*Contact{
		&Contact{"Dmitry", "+375290000000"},
		&Contact{"Alex", "+375290000010"},
		&Contact{"Ruslan", "+375290000015"},
	}

	fmt.Println("-------Before-------")
	Show(contacts)
	fmt.Println("-------After-------")
	Sort(contacts)

}

func Sort(contacts Contacts) {
	sort.Slice(contacts, func(i, j int) bool { return contacts[i].name < contacts[j].name })
	for _, contact := range contacts {
		fmt.Println(contact)
	}
}

func Show(list []*Contact) {
	for _, l := range list {
		fmt.Println(l)
	}
}
