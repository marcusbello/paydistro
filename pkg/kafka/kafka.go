package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

type Kafka struct {
	kafka *kafka.Conn
}

func (k *Kafka) Writer(ctx context.Context, key string, value map[string]interface{}) error {

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%s", key)),
		Value: []byte(fmt.Sprint(value)),
	}
	err := k.kafka.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}
	_, err = k.kafka.WriteMessages(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (k *Kafka) Reader() (*map[string]interface{}, error) {
	var resp map[string]interface{}

	resp = map[string]interface{}{
		"key-2384374346": "{some dict}",
	}

	return &resp, nil
}

func New(url, topic string) (*Kafka, error) {

	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", url, topic, partition)
	if err != nil {
		return nil, err
	}
	k := &Kafka{
		kafka: conn,
	}

	return k, nil
}
