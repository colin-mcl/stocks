package repo_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
)

func TestCreatePosition(t *testing.T) {
	// create position doesn't do any checking of values so there should not
	// be any error cases
	id, err := testRepo.CreatePosition(&models.Position{
		Symbol:        util.RandomString(50),
		HeldBy:        11,
		PurchasedAt:   time.Now(),
		PurchasePrice: rand.Float64() * 10000,
		Qty:           rand.Float64() * 100000,
	})

	assert.NoError(t, err)
	assert.Positive(t, id)
}
