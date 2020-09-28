package notify

import (
	"github.com/FernandoCagale/c4-notify/internal/errors"
	"github.com/FernandoCagale/c4-notify/pkg/entity"
)

type NotifyUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *NotifyUseCase {
	return &NotifyUseCase{
		repo: repo,
	}
}

func (usecase *NotifyUseCase) FindAll() (notify []*entity.Notify, err error) {
	return usecase.repo.FindAll()
}

func (usecase *NotifyUseCase) FindById(ID string) (notify *entity.Notify, err error) {
	return usecase.repo.FindById(ID)
}

func (usecase *NotifyUseCase) DeleteById(ID string) (err error) {
	return usecase.repo.DeleteById(ID)
}

func (usecase *NotifyUseCase) Create(customer *entity.Customer) error {
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
