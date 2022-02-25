package service

import (
	"address-book-go/config"
	"address-book-go/driver"
	"address-book-go/dto/apireq"
	sysAccRepo "address-book-go/internal/system/sys_account/repository"
	"address-book-go/pkg/valider"
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

func TestService_GenToken(t *testing.T) {
	// Arrange
	orm, _ := driver.NewXorm()
	sar := sysAccRepo.NewRepository(orm)
	ts := NewService(sar)

	// Has data
	req := apireq.GetSysAccountToken{
		Account:  "sys_account",
		Password: "A12345678",
	}

	// Act
	res, err := ts.GenToken(&req)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
