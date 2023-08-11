package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type MyHandler struct{}

func (h *MyHandler) HandleMessage(message *nsq.Message) error {
	log.Printf("\033[2K\r new message: %s", message.Body)
	message.Finish()
	return nil
}

func main() {

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("my-topic", "app-channel", config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Stop()

	consumer.AddHandler(&MyHandler{})
	err = consumer.ConnectToNSQD("localhost:4150")
	if err != nil {
		log.Fatal(err)
	}

	// two
	consumerTwo, err := nsq.NewConsumer("my-topic", "app-channel-two", config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumerTwo.Stop()

	consumerTwo.AddHandler(&MyHandler{})
	err = consumerTwo.ConnectToNSQD("localhost:4150")
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
