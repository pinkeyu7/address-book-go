package apires

import (
	"address-book-go/dto/model"
	"gopkg.in/guregu/null.v4"
)

type ListContact struct {
	List        []*model.Contact `json:"list"`
	Total       int              `json:"total"`
	CurrentPage int              `json:"current_page"`
	PerPage     int              `json:"per_page"`
	NextPage    null.Int         `json:"next_page" swaggertype:"string"`
}
