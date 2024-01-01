package producer

import (
	"log"
)

type Producer struct {
	Jobs chan int
}

func NewProducer(jobs chan int) Producer {
	return Producer{
		Jobs: jobs,
	}
}

func (p Producer) Produce(id int, results chan<- int) {
	for job := range p.Jobs {
		log.Printf("Worker %d completed job %d", id, job)
		// Simulating some work
		result := job * 2
		log.Printf("Worker %d completed job %d", id, job)
		results <- result
	}
}
