package apireq

type ListContact struct {
	AccountId int `form:"account_id" validate:"required"`
	Page      int `form:"page" validate:"required,numeric"`
	PerPage   int `form:"per_page" validate:"required,numeric"`
}

type GetContact struct {
	AccountId int `form:"account_id" validate:"required"`
}

type AddContact struct {
	AccountId int    `json:"account_id" validate:"required"`
	Name      string `json:"name" validate:"required,max=25"`
	Email     string `json:"email" validate:"required,max=40"`
	Phone     string `json:"phone" validate:"required,max=25"`
	Gender    *int   `json:"gender" validate:"required,oneof=0 1 2"`
}

type EditContact struct {
	AccountId int    `json:"account_id" validate:"required"`
	Name      string `json:"name" validate:"required,max=25"`
	Email     string `json:"email" validate:"required,max=40"`
	Phone     string `json:"phone" validate:"required,max=25"`
	Gender    *int   `json:"gender" validate:"required,oneof=0 1 2"`
}
