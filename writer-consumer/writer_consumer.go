package writerconsumer

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
)

// WriterConsumer struct
type WriterConsumer struct {
	q *nsq.Consumer
}

// New writer consumer creating
func New(q *nsq.Consumer) *WriterConsumer {
	return &WriterConsumer{
		q: q,
	}
}

// Init consumer write event
func (wc *WriterConsumer) Init(wg *sync.WaitGroup) {
	wc.q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %s", message.Body)
		log.Printf("Full response is: %v", message)

		addEvent(message)
		updateElastic(message)

		wg.Done()
		return nil
	}))
}

func addEvent(message *nsq.Message) {
	// TODO add event to couchbase bucket
}

func updateElastic(message *nsq.Message) {
	// TODO update elastic document
}
