package event

import (
	eventImp "github.com/FernandoCagale/c4-notify/internal/event"
	"github.com/FernandoCagale/c4-notify/pkg/domain/notify"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewNotify, notify.Set, eventImp.Set)
