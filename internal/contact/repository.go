package contact

import "address-book-go/dto/model"

type Repository interface {
	Find(offset, limit int) ([]*model.Contact, error)
	Count() (int, error)
}
