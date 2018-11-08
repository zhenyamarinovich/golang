package main

import (
	"fmt"
	"sync"
)

func output() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(output)
			wg.Done()
		}()
	}
	wg.Wait()
	//fmt.Scanln() // ждем ввода пользователя

}
