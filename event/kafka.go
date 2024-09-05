package event

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

const KafkaBroker = "localhost:9092"

var writer *kafka.Writer

func initKafkaWriter(topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{KafkaBroker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func SendMessage(topic string, msg []byte) error {
	// Lazily initialize Kafka writer if not already initialized
	if writer == nil {
		writer = initKafkaWriter(topic)
	}

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: msg,
		},
	)
	if err != nil {
		log.Println("could not write message:", err)
		return err
	}

	return nil
}

func CloseKafkaWriter() error {
	if writer != nil {
		if err := writer.Close(); err != nil {
			log.Println("could not close writer:", err)
			return err
		}
		writer = nil
	}
	return nil
}
