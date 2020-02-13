package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-notify/api/render"
	"github.com/FernandoCagale/c4-notify/internal/errors"
	"github.com/FernandoCagale/c4-notify/pkg/domain/notify"
	"github.com/FernandoCagale/c4-notify/pkg/entity"
	"github.com/gorilla/mux"
	"net/http"
)

type OrderHandler struct {
	usecase notify.UseCase
}

func NewOrder(usecase notify.UseCase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (handler *OrderHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	notifys, err := handler.usecase.FindAll()
	if err != nil {
		render.ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	render.Response(w, notifys, http.StatusOK)
}

func (handler *OrderHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	notify, err := handler.usecase.FindById(ID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, notify, http.StatusOK)
}

func (handler *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	err := handler.usecase.DeleteById(ID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var customer *entity.Customer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(customer); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		case errors.ErrConflict:
			render.ResponseError(w, err, http.StatusConflict)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusCreated)
}
