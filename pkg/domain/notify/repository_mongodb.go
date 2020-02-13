package notify

import (
	"github.com/FernandoCagale/c4-notify/internal/errors"
	"github.com/FernandoCagale/c4-notify/pkg/entity"
	"gopkg.in/mgo.v2"
)

const (
	COLLECTION = "notify"
	DATABASE   = "c4-notify-database"
)

type MongodbRepository struct {
	session *mgo.Session
}

func NewMongodbRepository(session *mgo.Session) *MongodbRepository {
	return &MongodbRepository{session}
}

func (repo *MongodbRepository) FindAll() (notify []*entity.Notify, err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.Find(nil).All(&notify)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return notify, nil
}

func (repo *MongodbRepository) FindById(ID string) (notify *entity.Notify, err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.FindId(ID).One(&notify)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return nil, errors.ErrNotFound
		default:
			return nil, errors.ErrInternalServer
		}
	}

	return notify, nil
}

func (repo *MongodbRepository) Create(notify *entity.Notify) (err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.Insert(notify)
	if err != nil {
		if mgo.IsDup(err) {
			return errors.ErrConflict
		}
		return errors.ErrInternalServer
	}
	return nil
}

func (repo *MongodbRepository) DeleteById(ID string) (err error) {
	coll := repo.session.DB(DATABASE).C(COLLECTION)

	err = coll.RemoveId(ID)
	if err != nil {
		switch err {
		case mgo.ErrNotFound:
			return errors.ErrNotFound
		default:
			return errors.ErrInternalServer
		}
	}

	return nil
}
