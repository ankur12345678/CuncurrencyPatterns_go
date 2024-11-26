package main

import (
	"fmt"
	"sync"
)

//In generator pattern then is a Provider and a Consumer
//Provider enters a value in the channel and consumer consumes

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//generator
		//it will add values to the channel (ONLY WHEN the prev value is consumed)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	wg.Add(1)
	go func() {
		//consumer
		for i := range ch {
			fmt.Printf("Message received by consumer is %v \n", i)
		}
		wg.Done()
	}()
	wg.Wait()
}

// NOTE: In the ouput it can be seen that it does step by step consumption

// OUTPUT:
// Message received by consumer is 1 
// Message received by consumer is 2 
// Message received by consumer is 3 
// Message received by consumer is 4 
// Message received by consumer is 5 
// Message received by consumer is 6 
// Message received by consumer is 7 
// Message received by consumer is 8 
// Message received by consumer is 9 
// Message received by consumer is 10 
