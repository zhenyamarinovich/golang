package main

import "fmt"

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

func (p person) printInfo() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Фамилия:", p.surname)
	fmt.Println("Возраст:", p.age)
	fmt.Println("Адрес:", fullAddress(p))
}

func fullAddress(p person) string {
	return p.addressInfo.country + " " + p.addressInfo.city + " " + p.addressInfo.street + " " + p.addressInfo.houseNumber
}
func main() {

	var Bob = person{
		name:    "Bob",
		surname: "Ivanov",
		age:     25,
		addressInfo: address{
			country:     "Belarus",
			city:        "Gomel",
			street:      "Grandma's street",
			houseNumber: "48a",
		},
	}
	Bob.printInfo()

}
