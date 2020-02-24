package routers

import (
	"github.com/FernandoCagale/c4-notify/api/event"
	"github.com/FernandoCagale/c4-notify/api/handlers"
	"github.com/gorilla/mux"
	"time"
)

type SystemRoutes struct {
	healthHandler *handlers.HealthHandler
	notifyHandler *handlers.OrderHandler
	notifyEvent   *event.NotifyEvent
}

func (routes *SystemRoutes) MakeEvents() {
	time.Sleep(5 * time.Second)

	go routes.notifyEvent.ProcessRegistered()
	go routes.notifyEvent.ProcessNotifyPayment()
	go routes.notifyEvent.ProcessNotifyOrder()
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/notify", routes.notifyHandler.Create).Methods("POST")
	r.HandleFunc("/notify", routes.notifyHandler.FindAll).Methods("GET")
	r.HandleFunc("/notify/{id}", routes.notifyHandler.FindById).Methods("GET")
	r.HandleFunc("/notify/{id}", routes.notifyHandler.DeleteById).Methods("DELETE")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, notifyHandler *handlers.OrderHandler, notifyEvent *event.NotifyEvent) *SystemRoutes {
	return &SystemRoutes{
		healthHandler: healthHandler,
		notifyHandler: notifyHandler,
		notifyEvent:   notifyEvent,
	}
}
