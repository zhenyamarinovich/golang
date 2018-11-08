package main

import "fmt"

func main() {
	f := func() {
		defer fmt.Println("World")
		fmt.Println("Hello")
	}
	f()
}
