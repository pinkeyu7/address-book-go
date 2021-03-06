package service

import (
	"address-book-go/config"
	"address-book-go/driver"
	"address-book-go/dto/apireq"
	sysAccRepo "address-book-go/internal/system/sys_account/repository"
	tokenRepo "address-book-go/internal/token/repository"
	"address-book-go/pkg/er"
	"address-book-go/pkg/valider"
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

func TestService_GenToken(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	rc, _ := driver.NewRedis()
	sar := sysAccRepo.NewRepository(orm)
	tc := tokenRepo.NewRedis(rc)
	ts := NewService(sar, tc)

	// No data
	req := apireq.GetSysAccountToken{
		Account:  "sys_account",
		Password: "123456",
	}

	// Act
	res, err := ts.GenToken(&req)

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, res)
	authErr := err.(*er.AppError)
	assert.Equal(t, http.StatusUnauthorized, authErr.StatusCode)
	assert.Equal(t, strconv.Itoa(er.UnauthorizedError), authErr.Code)

	// Has data
	req = apireq.GetSysAccountToken{
		Account:  "sys_account",
		Password: "A12345678",
	}

	// Act
	res, err = ts.GenToken(&req)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
