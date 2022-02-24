package seed

import (
	"address-book-go/dto/model"
	"address-book-go/internal/contact"
	"github.com/brianvoe/gofakeit/v4"
	"xorm.io/xorm"
)

func CreateContact(engine *xorm.Engine, name, email, phone string, gender int) error {
	con := model.Contact{
		Name:   name,
		Email:  email,
		Phone:  phone,
		Gender: &gender,
	}

	_, err := engine.Insert(&con)
	return err
}

func AllContact() []Seed {
	return []Seed{
		{
			Name: "Create Contact - 1",
			Run: func(engine *xorm.Engine) error {
				err := CreateContact(engine, gofakeit.Name(), gofakeit.Email(), gofakeit.Phone(), contact.Male)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "Create Contact - 2",
			Run: func(engine *xorm.Engine) error {
				err := CreateContact(engine, gofakeit.Name(), gofakeit.Email(), gofakeit.Phone(), contact.Male)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "Create Contact - 3",
			Run: func(engine *xorm.Engine) error {
				err := CreateContact(engine, gofakeit.Name(), gofakeit.Email(), gofakeit.Phone(), contact.Female)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "Create Contact - 4",
			Run: func(engine *xorm.Engine) error {
				err := CreateContact(engine, gofakeit.Name(), gofakeit.Email(), gofakeit.Phone(), contact.Secret)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
