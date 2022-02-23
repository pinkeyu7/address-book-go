package repository

import (
	"address-book-go/dto/model"
	"address-book-go/internal/contact"
)

type Repository struct {
}

func NewRepository() contact.Repository {
	return &Repository{}
}

func (r *Repository) Find(offset, limit int) ([]*model.Contact, error) {
	list := make([]*model.Contact, 0)

	// TODO - get from mysql
	list = append(list, &model.Contact{
		// Fake data
		Id:    1,
		Name:  "name_1",
		Phone: "phone_1",
		Email: "email_1",
	})
	list = append(list, &model.Contact{
		Id:    2,
		Name:  "name_2",
		Email: "email_2",
		Phone: "phone_2",
	})

	return list, nil
}

func (r *Repository) Count() (int, error) {
	// TODO - get from mysql
	count := 2
	return count, nil
}
