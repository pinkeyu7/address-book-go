package contact

import "address-book-go/dto/model"

type Repository interface {
	Insert(m *model.Contact) error
	Find(offset, limit int) ([]*model.Contact, error)
	FindOne(m *model.Contact) (*model.Contact, error)
	Count() (int, error)
	Update(m *model.Contact) error
}
