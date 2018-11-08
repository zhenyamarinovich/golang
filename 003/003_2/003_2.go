package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	printInfo() string
}

func output(a Printable) {
	a.printInfo()
}

type address struct {
	country     string
	city        string
	street      string
	houseNumber string
}

type person struct {
	name        string
	surname     string
	age         int
	addressInfo address
}

func (p person) printInfo() string {
	return p.name + " " + p.surname + " " +
		"Возраст: " + strconv.Itoa(p.age) + " " +
		"Адрес:" + p.addressInfo.printInfo()

}

func (a address) printInfo() string {
	return a.country + " " + a.city
}

func showPrintable(p []Printable) {
	for index, value := range p {
		fmt.Println(strconv.Itoa(index)+" | ", value.printInfo())
	}
}

func main() {

	var tmp = []Printable{
		person{
			name:    "Bob",
			surname: "Ivanov",
			age:     25,
			addressInfo: address{
				country:     "Belarus",
				city:        "Gomel",
				street:      "Grandma's street",
				houseNumber: "48a",
			},
		},
		person{
			name:    "Alice",
			surname: "Petrova",
			age:     22,
			addressInfo: address{
				country:     "Belarus",
				city:        "Gomel",
				street:      "Push street",
				houseNumber: "10",
			},
		},
		person{
			name:    "Dasha",
			surname: "Smirnova",
			age:     15,
			addressInfo: address{
				country:     "Belarus",
				city:        "Gomel",
				street:      "Simple street",
				houseNumber: "13b",
			},
		},
		address{
			country:     "Belarus",
			city:        "Gomel",
			street:      "Big street",
			houseNumber: "1a",
		},
	}

	showPrintable(tmp)
}
