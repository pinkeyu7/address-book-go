package repository

import (
	"address-book-go/config"
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

func TestRepository_Find(t *testing.T) {
	// Arrange
	cr := NewRepository()

	// Act
	testCases := []struct {
		Limit     int
		Offset    int
		WantCount int
	}{
		{
			5,
			0,
			2,
		},
		{
			5,
			0,
			2,
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

func TestRepository_Count(t *testing.T) {
	// Arrange
	cr := NewRepository()

	offset := 0
	limit := 10

	// Act
	data, err := cr.Find(offset, limit)

	// Assert
	assert.Nil(t, err)
	assert.Len(t, data, 2)
}
