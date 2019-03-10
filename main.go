package main

import (
	"log"
	"sync"

	"github.com/bitly/go-nsq"
	writerconsumer "github.com/vrgbrg/flowfluence/writer-consumer"
)

const (
	ConsumerTopic = "write_consumer"
	Host          = "127.0.0.1"
	Port          = "4150"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(ConsumerTopic, "ch", config)

	wc := writerconsumer.New(q)
	wc.Init(wg)

	err := q.ConnectToNSQD(Host + ":" + Port)
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
