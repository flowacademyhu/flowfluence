package writerconsumer

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_consumer", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", message)

		addEvent(message)
		updateElastic(message)

		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}

func addEvent(message *nsq.Message) {
	// TODO add event to couchbase bucket
}

func updateElastic(message *nsq.Message) {
	// TODO update elastic document
}
