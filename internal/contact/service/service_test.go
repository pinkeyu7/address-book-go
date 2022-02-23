package service

import (
	"address-book-go/config"
	"address-book-go/dto/apireq"
	contactRepo "address-book-go/internal/contact/repository"
	"address-book-go/pkg/valider"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
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
	cr := contactRepo.NewRepository()
	cs := NewService(cr)

	// Act
	testCases := []struct {
		PerPage   int
		Page      int
		WantCount int
	}{
		{
			5,
			1,
			2,
		},
		{
			5,
			1,
			2,
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
