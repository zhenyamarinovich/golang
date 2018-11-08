package main

import (
	"fmt"
	"strconv"
)

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
		"Адрес: " + fullAddress(p) + " "
}

func fullAddress(p person) string {
	return p.addressInfo.country + " " + p.addressInfo.city + " " + p.addressInfo.street + " " + p.addressInfo.houseNumber
}
func main() {

	var people = [3]person{
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
	}

	for index, value := range people {
		fmt.Println(strconv.Itoa(index)+" | ", value.printInfo())
	}

}
