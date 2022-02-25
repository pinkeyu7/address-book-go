package seed

import (
	"address-book-go/dto/model"
	"github.com/brianvoe/gofakeit/v4"
	"xorm.io/xorm"
)

func CreateSysAccount(engine *xorm.Engine, account, name, email, phone, password string) error {
	con := model.SysAccount{
		Account:  account,
		Phone:    phone,
		Email:    email,
		Password: password,
		Name:     name,
	}

	_, err := engine.Insert(&con)
	return err
}

func AllSysAccount() []Seed {
	return []Seed{
		{
			Name: "Create System Account - 1",
			Run: func(engine *xorm.Engine) error {
				err := CreateSysAccount(engine, "sys_account", gofakeit.Name(), gofakeit.Email(), gofakeit.Phone(), "123456")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
