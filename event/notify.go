package event

import (
	"github.com/FernandoCagale/c4-notify/internal/broker/consumer"
)

type NotifyEvent struct {
	consumer consumer.Consumer
}

func NewNotify(consumer consumer.Consumer) *NotifyEvent {
	return &NotifyEvent{
		consumer: consumer,
	}
}

func (event *NotifyEvent) MakeEvents() {
	go event.consumer.Customer()
	go event.consumer.Order()
	go event.consumer.Payment()
}
