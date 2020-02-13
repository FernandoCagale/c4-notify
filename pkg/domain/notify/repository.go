package notify

import "github.com/FernandoCagale/c4-notify/pkg/entity"

type Repository interface {
	Create(notify *entity.Notify) (err error)
	FindAll() (notify []*entity.Notify, err error)
	FindById(ID string) (notify *entity.Notify, err error)
	DeleteById(ID string) (err error)
}
