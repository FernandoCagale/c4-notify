package notify

import (
	"github.com/FernandoCagale/c4-notify/internal/errors"
	"github.com/FernandoCagale/c4-notify/pkg/entity"
)

type InMemoryRepository struct {
	m map[string]*entity.Notify
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{map[string]*entity.Notify{}}
}

func (repo *InMemoryRepository) FindAll() (notifys []*entity.Notify, err error) {
	for _, notify := range repo.m {
		notifys = append(notifys, notify)
	}
	return notifys, nil
}

func (repo *InMemoryRepository)  FindById(ID string) (notify *entity.Notify, err error) {
	for _, notify := range repo.m {
		if notify.Code == ID {
			return notify, nil
		}
	}
	return nil, errors.ErrNotFound
}

func (repo *InMemoryRepository)  DeleteById(ID string) (err error) {
	for _, notify := range repo.m {
		if notify.Code == ID {
			delete(repo.m, ID)
			return nil
		}
	}
	return errors.ErrNotFound
}

func (repo *InMemoryRepository) Create(e *entity.Notify) (err error) {
	notify := repo.m[e.Code]

	if notify == nil {
		repo.m[e.Code] = e
		return nil
	}

	return nil
}
