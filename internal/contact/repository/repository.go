package repository

import (
	"address-book-go/dto/model"
	"address-book-go/internal/contact"
	"xorm.io/xorm"
)

type Repository struct {
	orm *xorm.EngineGroup
}

func NewRepository(orm *xorm.EngineGroup) contact.Repository {
	return &Repository{
		orm: orm,
	}
}

func (r *Repository) Insert(m *model.Contact) error {
	_, err := r.orm.Insert(m)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Find(offset, limit int) ([]*model.Contact, error) {
	list := make([]*model.Contact, 0)

	err := r.orm.Limit(limit, offset).Find(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repository) FindOne(m *model.Contact) (*model.Contact, error) {
	has, err := r.orm.Get(m)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return m, nil
}

func (r *Repository) Count() (int, error) {
	count, err := r.orm.Count(&model.Contact{})
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *Repository) Update(m *model.Contact) error {
	_, err := r.orm.ID(m.Id).Update(m)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(contactId int) error {
	_, err := r.orm.ID(contactId).Delete(&model.Contact{})
	if err != nil {
		return err
	}

	return nil
}
