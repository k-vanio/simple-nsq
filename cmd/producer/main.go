package main

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
)

func main() {

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("localhost:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Stop()

	topic := "my-topic"

	for i := 0; i < 3000; i++ {
		message := []byte(fmt.Sprintf("value: %v", i))
		err = producer.Publish(topic, message)
		if err != nil {
			log.Fatal(err)
		}
	}

}
