// microservice1.go

package main

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	brokerList := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokerList, nil)
	if err != nil {
		log.Fatalln("Failed to create producer:", err)
	}
	defer producer.Close()

	consumer, err := sarama.NewConsumer(brokerList, nil)
	if err != nil {
		log.Fatalln("Failed to create consumer:", err)
	}
	defer consumer.Close()

	topic := "my_topic"

	// Data sending loop for Microservice 1
	go func() {
		for i := 0; i < 10; i++ {
			data := []byte("Message from Microservice 1: " + string(i))
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.ByteEncoder(data),
			}
			_, _, err := producer.SendMessage(msg)
			if err != nil {
				log.Println("Failed to send message:", err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Data receiving loop for Microservice 1
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to create partition consumer:", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Microservice 1 Received message: %s\n", string(msg.Value))
		}
	}
}
