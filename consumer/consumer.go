package consumer

import (
	"log"
)

type Consumer struct {
	Results chan int
}

func NewConsumer(results chan int) Consumer {
	return Consumer{
		Results: results,
	}
}

func (c Consumer) Consume() {
	for r := range c.Results {
		log.Printf("Result: %d", r)
	}
}
