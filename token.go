package main

import (
	"flag"
	"log"

	"github.com/marcusbello/paydistro/internal/token/server"
)

var addr = flag.String("addr", "127.0.0.1:6655", "The address to run on.")
var kafkaURL = flag.String("kafkaURL", "127.0.0.1:9092", "The address on which kafka cluster is running")
var kafkaTopic = flag.String("kafkaTopic", "buy_token", "Kafka topic")
var kafkaGroupID = flag.String("kafkaGroupID", "getuser", "Kafka group")

func main() {
	flag.Parse()
	s, err := server.New(*addr, *kafkaURL, *kafkaTopic, *kafkaGroupID)
	if err != nil {
		panic(err)
	}
	done := make(chan error, 1)

	log.Println("Starting server at: ", *addr)
	go func() {
		defer close(done)
		done <- s.Start()
	}()

	err = <-done
	log.Println("Server exited with error: ", err)
}
