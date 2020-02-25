package pkg

import (
	"github.com/FernandoCagale/c4-notify/api/handlers"
	"github.com/FernandoCagale/c4-notify/api/routers"
	"github.com/FernandoCagale/c4-notify/internal/event"
	"github.com/FernandoCagale/c4-notify/pkg/domain/notify"
	"github.com/google/wire"
)

var Container = wire.NewSet(notify.Set, handlers.Set, routers.Set, event.Set)
