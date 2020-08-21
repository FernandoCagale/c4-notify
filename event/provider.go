package event

import (
	"github.com/FernandoCagale/c4-notify/internal/broker/consumer"
	"github.com/FernandoCagale/c4-notify/pkg/domain/notify"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewNotify, notify.Set, consumer.Set)
