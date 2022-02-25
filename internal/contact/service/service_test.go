package service

import (
	"address-book-go/config"
	"address-book-go/driver"
	"address-book-go/dto/apireq"
	"address-book-go/dto/model"
	"address-book-go/internal/contact"
	contactRepo "address-book-go/internal/contact/repository"
	"address-book-go/pkg/er"
	"address-book-go/pkg/valider"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")
	if remoteBranch == "" {
		// load env
		err := godotenv.Load(config.GetBasePath() + "/.env")
		if err != nil {
			log.Panicln(err)
		}
	}

	valider.Init()
}

func TestService_List(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := contactRepo.NewRepository(orm)
	cs := NewService(cr)

	// Act
	testCases := []struct {
		PerPage   int
		Page      int
		WantCount int
	}{
		{
			2,
			1,
			2,
		},
		{
			10,
			1,
			4,
		},
		{
			10,
			2,
			0,
		},
	}
	// Act
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("List Contact,Page:%d,PerPage:%d", tc.Page, tc.PerPage), func(t *testing.T) {
			req := apireq.ListContact{
				Page:    tc.Page,
				PerPage: tc.PerPage,
			}
			res, err := cs.List(&req)
			assert.Nil(t, err)
			assert.Len(t, res.List, tc.WantCount)
		})
	}
}

func TestService_Get(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := contactRepo.NewRepository(orm)
	cs := NewService(cr)

	// No data
	contactId := 10

	// Act
	con, err := cs.Get(contactId)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, con)
	notFoundErr := err.(*er.AppError)
	assert.Equal(t, http.StatusBadRequest, notFoundErr.StatusCode)
	assert.Equal(t, strconv.Itoa(er.ResourceNotFoundError), notFoundErr.Code)

	// Has data
	contactId = 2

	// Act
	con, err = cs.Get(contactId)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, contactId, con.Id)
}

func TestService_Add(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := contactRepo.NewRepository(orm)
	cs := NewService(cr)

	gender := contact.Male
	req := apireq.AddContact{
		Name:   "test_name",
		Email:  "test_email",
		Phone:  "test_phone",
		Gender: &gender,
	}

	// Act
	err := cs.Add(&req)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.Where("name = ?", req.Name).Delete(&model.Contact{})
}

func TestService_Edit(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := contactRepo.NewRepository(orm)
	cs := NewService(cr)

	contactId := 1
	con := model.Contact{Id: contactId}
	_, _ = orm.Get(&con)

	gender := contact.Male
	req := apireq.EditContact{
		Name:   "test_name",
		Email:  "test_email",
		Phone:  "test_phone",
		Gender: &gender,
	}

	// Act
	err := cs.Edit(contactId, &req)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.ID(contactId).Update(&con)
}
