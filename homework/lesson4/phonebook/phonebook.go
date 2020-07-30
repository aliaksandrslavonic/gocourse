package main

import "fmt"

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

	Show(contacts)
	Sort(contacts)

}

func Sort(contacts Contacts) {
	for idx, contact := range contacts {
		fmt.Println(idx)
		fmt.Println(contact)
		// if contact.name > contact[idx+1].name {
		// 	fmt.Println()
		// }
	}
}

func Show(list []*Contact) {
	for _, l := range list {
		fmt.Println(l)
	}
}
