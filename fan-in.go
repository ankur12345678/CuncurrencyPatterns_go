package main

import (
	"fmt"
	"sync"
)

//In FAN-IN pattern there is multiple producer , a consumer , a channel
//multiple producer produces value in the channel and the consumer consumes value from the channel

func producer(input []int, ch chan int, wg *sync.WaitGroup) {
	go func() {
		for _, val := range input {
			ch <- val
		}
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	data1 := []int{1, 2, 3, 4, 5}
	data2 := []int{6, 7, 8, 9, 10}

	//common channel
	ch := make(chan int)

	//fire up 2 producers and share the access of channel with them and they will constantly add data in the channel
	wg.Add(2)
	producer(data1, ch, &wg)
	producer(data2, ch, &wg)

	//consumer
	go func() {
		for value := range ch {
			fmt.Printf("Message received by consumer is %v \n", value)
		}
	}()
	wg.Wait()
	close(ch)

}

//OUTPUT:
// Message received by consumer is 1 
// Message received by consumer is 2 
// Message received by consumer is 6 
// Message received by consumer is 7 
// Message received by consumer is 8 
// Message received by consumer is 3 
// Message received by consumer is 9 
// Message received by consumer is 4 
// Message received by consumer is 10 
// Message received by consumer is 5 