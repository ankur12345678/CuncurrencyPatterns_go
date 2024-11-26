package main

import (
	"fmt"
	"sync"
)

//In PIPELINE pattern we will do some processing on data and pass to a new channel
//eg processing --chan1--> processing2 --chan2--> processing3 --chan3--> so on
//the pattern is used if we have to MULTIPLE processing on some data

//each func takes a channel read data (and transform it..) and send all the data to a new channel and return the new channel as return type

// eg:
// FILTER
//
//	|
//	V
//
// SQUARE
//
//	 |
//	 V
//	HALF
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	ch1 := filter(ch)
	ch2 := square(ch1)
	ch3 := half(ch2)
	wg.Add(1)
	go func() {
		for value := range ch3 {
			fmt.Printf("Received : %v\n", value)
		}
		wg.Done()
	}()
	wg.Wait()
}

func filter(ch chan int) chan int {
	ch1 := make(chan int)
	go func() {
		for val := range ch {
			if val%2 == 0 {
				ch1 <- val
			}
		}
		close(ch1)
	}()
	return ch1
}

func square(ch1 chan int) chan int {
	ch2 := make(chan int)
	go func() {
		for val := range ch1 {

			ch2 <- val * val

		}
		close(ch2)
	}()
	return ch2
}

func half(ch2 chan int) chan int {
	ch3 := make(chan int)
	go func() {
		for val := range ch2 {

			ch3 <- val / 2

		}
		close(ch3)
	}()
	return ch3
}


//OUTPUT:
// Received : 2
// Received : 8
// Received : 18
// Received : 32
// Received : 50