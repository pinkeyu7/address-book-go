package service

import (
	"address-book-go/dto/apireq"
	"address-book-go/dto/apires"
	"address-book-go/dto/model"
	"address-book-go/internal/contact"
	"address-book-go/pkg/er"
	"gopkg.in/guregu/null.v4"
	"net/http"
)

type Service struct {
	contactRepo contact.Repository
}

func NewService(cr contact.Repository) contact.Service {
	return &Service{
		contactRepo: cr,
	}
}

func (s *Service) List(req *apireq.ListContact) (*apires.ListContact, error) {
	page := req.Page
	perPage := req.PerPage

	if page <= 1 {
		page = 1
	}

	if perPage <= 1 {
		perPage = 1
	}

	offset := (page - 1) * perPage

	total, err := s.contactRepo.Count()
	if err != nil {
		findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "count contact error.", err)
		return nil, findErr
	}

	data, err := s.contactRepo.Find(offset, perPage)
	if err != nil {
		findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "find contact error.", err)
		return nil, findErr
	}

	res := &apires.ListContact{
		List:        data,
		Total:       total,
		CurrentPage: page,
		PerPage:     perPage,
	}

	// 判斷 offset 加上資料筆數，是否仍小於總筆數,是的話回傳下一頁頁數
	dataCount := len(data)
	if (offset + dataCount) < total {
		res.NextPage = null.IntFrom(int64(page) + int64(1))
	}

	return res, nil
}

func (s *Service) Get(contactId int) (*model.Contact, error) {
	return nil, nil
}

func (s *Service) Add(req *apireq.AddContact) error {
	return nil
}

func (s *Service) Edit(contactId int, req *apireq.EditContact) error {
	return nil
}