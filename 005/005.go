package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func request(countError *int, needTime int, address string, wg *sync.WaitGroup, tmp *time.Duration) {

	tBegin := time.Now()
	client := http.Client{
		Timeout: time.Duration(needTime) * time.Second,
	}

	resp, err := client.Get(address)
	if err != nil { //счетчик
		fmt.Println(err)
		*countError++
		return
	}
	defer resp.Body.Close()
	wg.Done()
	tFinish := time.Now()
	*tmp = tFinish.Sub(tBegin)

}

func main() {

	var address string
	flag.StringVar(&address, "address", "https://google.com", "a string var")
	var count int
	flag.IntVar(&count, "count", 6, "an int")
	var needTime int
	flag.IntVar(&needTime, "time", 4, "an int")
	var countError int
	var wg sync.WaitGroup

	t0 := time.Now()
	var mass []time.Duration
	var tmp time.Duration

	wg.Add(count)
	for i := 0; i < count; i++ {
		go request(&countError, needTime, address, &wg, &tmp)
		mass = append(mass, tmp)
	}
	wg.Wait()
	timeFinish := time.Now()
	tMin := mass[0]
	tMax := mass[0]
	for _, value := range mass {
		fmt.Println(value)
	}
	for i := 0; i < len(mass); i++ {
		if tMin > mass[i] {
			tMin = mass[i]
		}
		if tMax < mass[i] {
			tMax = mass[i]
		}
	}
	fmt.Println("Время, за которое отработали все запросы ", timeFinish.Sub(t0))
	fmt.Println("Кол-во ответов, которых не дождались", countError)
	fmt.Println("Максимальное время на запрос", tMax)
	fmt.Println("Минимальное время на запрос", tMin)
}
