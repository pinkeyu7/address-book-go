package repository

import (
	"address-book-go/config"
	"address-book-go/driver"
	"address-book-go/dto/model"
	"address-book-go/internal/contact"
	"address-book-go/pkg/valider"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
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

func TestRepository_Insert(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := NewRepository(orm)

	gender := contact.Male
	m := model.Contact{
		Name:   "test_name",
		Email:  "test_email",
		Phone:  "test_phone",
		Gender: &gender,
	}

	// Act
	err := cr.Insert(&m)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.ID(m.Id).Delete(&model.Contact{})
}

func TestRepository_Find(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := NewRepository(orm)

	// Act
	testCases := []struct {
		Limit     int
		Offset    int
		WantCount int
	}{
		{
			2,
			0,
			2,
		},
		{
			10,
			0,
			4,
		},
	}
	// Act
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("Find Contact,Offset:%d,Limit:%d", tc.Offset, tc.Limit), func(t *testing.T) {
			data, err := cr.Find(tc.Offset, tc.Limit)
			assert.Nil(t, err)
			assert.Len(t, data, tc.WantCount)
		})
	}
}

func TestRepository_FindOne(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := NewRepository(orm)

	// No data
	// Act
	res, err := cr.FindOne(&model.Contact{Id: 100})

	// Assert
	assert.Nil(t, err)
	assert.Nil(t, res)

	// Has data
	// Act
	res, err = cr.FindOne(&model.Contact{Id: 1})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 1, res.Id)
}

func TestRepository_Count(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := NewRepository(orm)

	// Act
	count, err := cr.Count()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 4, count)
}

func TestRepository_Update(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	cr := NewRepository(orm)

	con := model.Contact{Id: 1}
	_, _ = orm.Get(&con)

	gender := contact.Male
	m := model.Contact{
		Id:     1,
		Name:   "test_name",
		Email:  "test_email",
		Phone:  "test_phone",
		Gender: &gender,
	}

	// Act
	err := cr.Update(&m)

	// Assert
	assert.Nil(t, err)

	// Teardown
	_, _ = orm.ID(con.Id).Update(&con)
}
