package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// Connect to NATS
	brokerUrl := "nats://127.0.0.1:4222"
	opts := []nats.Option{
		nats.Name("NATS Publisher"),
		nats.Compression(true),
	}
	conn, err := nats.Connect(brokerUrl, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queue := make(chan *nats.Msg)
	sub, err := conn.ChanSubscribe("my-test", queue)
	if err != nil {
		panic(err)
	}

	fmt.Println("Waiting for messages!")
	for m := range queue {
		fmt.Printf("received: %#v\n", m)
		if err := m.Ack(); err != nil {
			panic(err)
		}
	}
	if err := sub.Unsubscribe(); err != nil {
		panic(err)
	}
}
