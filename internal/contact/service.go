package contact

import (
	"address-book-go/dto/apireq"
	"address-book-go/dto/apires"
	"address-book-go/dto/model"
)

type Service interface {
	List(req *apireq.ListContact) (*apires.ListContact, error)
	Get(contactId int) (*model.Contact, error)
	Add(req *apireq.AddContact) error
	Edit(contactId int, req *apireq.EditContact) error
}
