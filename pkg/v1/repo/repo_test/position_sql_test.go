package repo_test

import (
	"database/sql"
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

func TestGetPosition(t *testing.T) {
	// should not exist
	p, err := testRepo.GetPosition(-1)

	assert.Error(t, err)
	assert.Nil(t, p)
	assert.EqualError(t, err, sql.ErrNoRows.Error())

	p, err = testRepo.GetPosition(0)

	assert.Error(t, err)
	assert.Nil(t, p)
	assert.EqualError(t, err, sql.ErrNoRows.Error())

	pos := &models.Position{
		Symbol:        "AAPL",
		HeldBy:        rand.Intn(10000) * 50,
		PurchasedAt:   time.Now(),
		PurchasePrice: rand.Float64() * 500,
		Qty:           rand.Float64() * 500,
	}

	id, err := testRepo.CreatePosition(pos)
	assert.NoError(t, err)

	p, err = testRepo.GetPosition(id)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.ID, id)
	assert.Equal(t, pos.Symbol, p.Symbol)
	assert.Equal(t, pos.HeldBy, p.HeldBy)
	assert.WithinDuration(t, pos.PurchasedAt, p.PurchasedAt, time.Minute)
	assert.Equal(t, pos.PurchasePrice, p.PurchasePrice)
	assert.Equal(t, pos.Qty, p.Qty)
}

func TestGetPositions(t *testing.T) {
	ps, err := testRepo.GetPositions("fake symbol", 11)
	assert.NoError(t, err)
	assert.Nil(t, ps)

	ps, err = testRepo.GetPositions("AAPL", 1)
	assert.Nil(t, ps)

	ps, err = testRepo.GetPositions("TSLA", 11)
	assert.NotNil(t, ps)
	assert.NoError(t, err)
	assert.Equal(t, 11, len(ps))

	for _, p := range ps {
		assert.Equal(t, 11, p.HeldBy)
		assert.Equal(t, 210.1, p.PurchasePrice)
		assert.Equal(t, 2.5, p.Qty)
		assert.Equal(t, "TSLA", p.Symbol)
		assert.Positive(t, p.ID)
	}
}

func TestGetPortfolio(t *testing.T) {
	ps, err := testRepo.GetPortfolio(1)
	assert.Nil(t, ps)
	assert.NoError(t, err)

	ps, err = testRepo.GetPortfolio(11)
	assert.NotNil(t, ps)
	assert.NoError(t, err)

	for _, p := range ps {
		assert.True(t, (p.Symbol == "TSLA" || p.Symbol == "AAPL"))
		assert.GreaterOrEqual(t, p.PurchasePrice, 210.1)
		assert.GreaterOrEqual(t, p.Qty, 2.5)
		assert.Equal(t, 11, p.HeldBy)
	}
}
