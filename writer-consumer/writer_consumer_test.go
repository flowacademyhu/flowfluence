package writerconsumer

import (
	"log"
	"testing"

	"github.com/bitly/go-nsq"
)

const (
	consumerTopic = "write_consumer"
	host          = "127.0.0.1"
	port          = "4150"
)

func TestWriterConsumer(t *testing.T) {

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(host+":"+port, config)

	err := w.Publish(consumerTopic, []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	log.Println("Message sent!")

	w.Stop()

}
