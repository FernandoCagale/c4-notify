package notify

import (
	"github.com/FernandoCagale/c4-notify/internal/errors"
	"github.com/FernandoCagale/c4-notify/internal/event"
	"github.com/FernandoCagale/c4-notify/pkg/entity"
)

type OrderUseCase struct {
	repo  Repository
	event event.Event
}

func NewUseCase(repo Repository, event event.Event) *OrderUseCase {
	return &OrderUseCase{
		repo:  repo,
		event: event,
	}
}

func (usecase *OrderUseCase) FindAll() (notify []*entity.Notify, err error) {
	return usecase.repo.FindAll()
}

func (usecase *OrderUseCase) FindById(ID string) (notify *entity.Notify, err error) {
	return usecase.repo.FindById(ID)
}

func (usecase *OrderUseCase) DeleteById(ID string) (err error) {
	return usecase.repo.DeleteById(ID)
}


func (usecase *OrderUseCase) Create(customer *entity.Customer) error {
	err := customer.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	notify := customer.ToNotify()

	if err = usecase.repo.Create(notify); err != nil {
		return err
	}

	return nil
}
