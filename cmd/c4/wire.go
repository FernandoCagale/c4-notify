//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-notify/api/routers"
	"github.com/FernandoCagale/c4-notify/event"
	"github.com/FernandoCagale/c4-notify/internal/datastore"
	"github.com/FernandoCagale/c4-notify/pkg"
	"github.com/google/wire"
	"gopkg.in/mgo.v2"
)

func SetupApplication(session *mgo.Session) (*routers.SystemRoutes, error) {
	wire.Build(pkg.Container)
	return nil, nil
}

func SetupMongoDB() (*mgo.Session, error) {
	wire.Build(datastore.Set)
	return nil, nil
}

func SetupEvents(session *mgo.Session) (*event.NotifyEvent, error) {
	wire.Build(event.Set)
	return nil, nil
}