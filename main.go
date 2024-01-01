package main

import (
	"sync"

	"github.com/diegom0ta/broker/consumer"
	"github.com/diegom0ta/broker/producer"
)

const (
	numJobs    = 10
	numWorkers = 3
)

func main() {

	send := make(chan int, numJobs)
	rcv := make(chan int, numJobs)

	p := producer.NewProducer(send)
	c := consumer.NewConsumer(rcv)

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			p.Produce(workerID, c.Results)
			wg.Done()
		}(w)
	}

	for j := 1; j <= numJobs; j++ {
		p.Jobs <- j
	}
	close(p.Jobs)

	wg.Wait()

	close(c.Results)
	c.Consume()
}
