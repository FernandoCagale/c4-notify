package event

import (
	"encoding/json"
	"fmt"
	"github.com/FernandoCagale/c4-notify/internal/event"
	"github.com/FernandoCagale/c4-notify/pkg/domain/notify"
	"github.com/FernandoCagale/c4-notify/pkg/entity"
	"log"
)

const (
	TOPIC = "topic.registered"
	QUEUE = "notify"
)

type NotifyEvent struct {
	usecase notify.UseCase
	event   event.Event
}

func NewNotify(usecase notify.UseCase, event event.Event) *NotifyEvent {
	return &NotifyEvent{
		usecase: usecase,
		event:   event,
	}
}

func (eventNotify *NotifyEvent) ProcessRegistered() {
	messages, err := eventNotify.event.Subscribe(TOPIC, QUEUE)
	if err != nil {
		fmt.Println(err.Error())
	}

	for msg := range messages {
		log.Printf("received message: %s, CUSTOMER: %s", msg.UUID, string(msg.Payload))

		var customer entity.Customer

		if err := json.Unmarshal(msg.Payload, &customer); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		if err = eventNotify.usecase.Create(&customer); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		msg.Ack() //TODO x-dead-letter-exchange
	}
}

func (eventNotify *NotifyEvent) ProcessNotify() {
	messages, err := eventNotify.event.Subscribe(TOPIC, QUEUE)
	if err != nil {
		fmt.Println(err.Error())
	}

	for msg := range messages {
		log.Printf("received message: %s, NOTIFY: %s", msg.UUID, string(msg.Payload))

		var customer entity.Customer

		if err := json.Unmarshal(msg.Payload, &customer); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		if err = eventNotify.usecase.Create(&customer); err != nil {
			fmt.Println(err.Error())
			msg.Nacked()
		}

		msg.Ack() //TODO x-dead-letter-exchange
	}
}
