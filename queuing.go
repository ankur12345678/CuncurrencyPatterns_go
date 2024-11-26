package main

import (
	"fmt"
	"sync"
)

//in queuing pattern, workers can pick 'n' jobs from the queue and work at it on one time

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 2)
	workerLimit := 4 //worker limit is required for worker-pool pattern
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < workerLimit; i++ {
		wg.Add(1)
		worker(ch, i+1, &wg)
	}
	wg.Wait()
}

func worker(ch chan int, num int, wg *sync.WaitGroup) {
	go func() {
		for value := range ch {
			fmt.Printf("worker %v consumed task %v\n", num, value)
		}
		wg.Done()
	}()
}


//OUTPUT:
// worker 2 consumed task 1
// worker 2 consumed task 5
// worker 3 consumed task 3
// worker 3 consumed task 7
// worker 3 consumed task 8
// worker 3 consumed task 9
// worker 4 consumed task 2
// worker 2 consumed task 6
// worker 3 consumed task 10
// worker 1 consumed task 4
