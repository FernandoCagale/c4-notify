package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FernandoCagale/c4-notify/pkg/domain/notify"
	"github.com/FernandoCagale/c4-notify/pkg/entity"
	"github.com/segmentio/kafka-go"
	"os"
	"strings"
)

const (
	CUSTOMER_TOPIC = "customer.registered"
	ORDER_TOPIC    = "order.registered"
	PAYMENT_TOPIC  = "payment.registered"
)

type ConsumerKafka struct {
	address string
	usecase notify.UseCase
}

func New(usecase notify.UseCase, ) *ConsumerKafka {
	return &ConsumerKafka{
		usecase: usecase,
		address: os.Getenv("ADDRESS_KAFKA"),
	}
}

func (e *ConsumerKafka) Customer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(e.address, ","),
		Topic:    CUSTOMER_TOPIC,
		GroupID:  "c4-notify",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err.Error()) //TODO
		}

		var customer entity.Customer

		if err := json.Unmarshal(m.Value, &customer); err != nil {
			fmt.Println(err.Error()) //TODO
		}

		if err := e.usecase.Create(&customer); err != nil {
			fmt.Println(err.Error()) //TODO
		}

		fmt.Println("--------------CUSTOMER-------------")
		fmt.Println(string(m.Value))
		fmt.Println("--------------------------------")

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func (e *ConsumerKafka) Order() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(e.address, ","),
		Topic:    ORDER_TOPIC,
		GroupID:  "c4-notify",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err.Error()) //TODO
		}

		fmt.Println("--------------ORDER-------------")
		fmt.Println(string(m.Value))
		fmt.Println("--------------------------------")

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func (e *ConsumerKafka) Payment() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(e.address, ","),
		Topic:    PAYMENT_TOPIC,
		GroupID:  "c4-notify",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer reader.Close()

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println(err.Error()) //TODO
		}

		fmt.Println("--------------PAYMENT-------------")
		fmt.Println(string(m.Value))
		fmt.Println("--------------------------------")

		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
