package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main(){
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}

	// Simple Publisher
	err := nc.Publish("foo", []byte("Hello World"))
	if err == nil {
		log.Println("Message published")
	}

	defer nc.Close()
}