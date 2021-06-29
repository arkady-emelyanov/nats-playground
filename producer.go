package main

import (
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

	// publish one message
	if err := conn.PublishRequest("my-test", "with-ok", []byte("hello world")); err != nil {
		panic(err)
	}
	if err := conn.Flush(); err != nil {
		panic(err)
	}
	if err := conn.LastError(); err != nil {
		panic(err)
	}
}
