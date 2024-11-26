package main

import (
	"fmt"
	"sync"
)

//In FAN-OUT pattern there is only one producer and multiple consumers and a channel
//producer produces on the channel and all the consumers can consume it
// if we define a fixed number of consumers(limit) then fanout pattern becomes worker-pool pattern

func consumer(ch chan int, num int, wg *sync.WaitGroup) {
	go func() {
		for val := range ch {
			fmt.Printf("consumer %v : consumed value is %v \n", num, val)
		}
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	go func() {
		//producer
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	wg.Add(2)
	consumer(ch, 1, &wg)
	consumer(ch, 2, &wg)
	wg.Wait()
}

//OUTPUT:
// consumer 2 : consumed value is 1 
// consumer 2 : consumed value is 3 
// consumer 2 : consumed value is 4 
// consumer 2 : consumed value is 5 
// consumer 2 : consumed value is 6 
// consumer 2 : consumed value is 7 
// consumer 2 : consumed value is 8 
// consumer 2 : consumed value is 9 
// consumer 1 : consumed value is 2 
// consumer 2 : consumed value is 10 