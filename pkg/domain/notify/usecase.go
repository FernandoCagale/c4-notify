package notify

import (
	"github.com/FernandoCagale/c4-notify/pkg/entity"
)

type UseCase interface {
	Create(customer *entity.Customer) (err error)
	FindAll() (notify []*entity.Notify, err error)
	FindById(ID string) (notify *entity.Notify, err error)
	DeleteById(ID string) (err error)
}
