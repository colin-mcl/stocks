package usecase_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pkg/v1/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCreatePosition(t *testing.T) {
	id, err := testUC.CreatePosition(&models.Position{
		Symbol:        "GOOG",
		HeldBy:        rand.Intn(1000) * 50,
		PurchasedAt:   time.Now(),
		PurchasePrice: rand.Float64() * 1000,
		Qty:           rand.Float64() * 10000,
	})

	assert.NoError(t, err)
	assert.Positive(t, id)
}

func TestGetPosition(t *testing.T) {
	p, err := testUC.GetPosition(-1)
	assert.Error(t, err)
	assert.Nil(t, p)
	assert.EqualError(t, err, usecase.ErrDoesNotExist.Error())

	p, err = testUC.GetPosition(0)
	assert.Error(t, err)
	assert.Nil(t, p)
	assert.EqualError(t, err, usecase.ErrDoesNotExist.Error())

	p, err = testUC.GetPosition(1200000)
	assert.Error(t, err)
	assert.Nil(t, p)
	assert.EqualError(t, err, usecase.ErrDoesNotExist.Error())

	p, err = testUC.GetPosition(1)
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "TSLA", p.Symbol)
	assert.Equal(t, 11, p.HeldBy)
	assert.Equal(t, 210.1, p.PurchasePrice)
	assert.True(t, time.Now().After(p.PurchasedAt))
	assert.Equal(t, 2.5, p.Qty)

	p = &models.Position{
		Symbol:        "GOOG",
		HeldBy:        rand.Intn(10000),
		PurchasedAt:   time.Now(),
		PurchasePrice: rand.Float64() * 1000,
		Qty:           rand.Float64() * 10000,
	}

	id, err := testUC.CreatePosition(p)
	assert.NoError(t, err)

	res, err := testUC.GetPosition(id)
	assert.NoError(t, err)
	assert.Equal(t, p.Symbol, res.Symbol)
	assert.Equal(t, p.HeldBy, res.HeldBy)
	assert.Equal(t, p.PurchasePrice, res.PurchasePrice)
	assert.Equal(t, p.Qty, res.Qty)
	assert.WithinDuration(t, p.PurchasedAt, res.PurchasedAt, time.Second)
}

func TestGetPositions(t *testing.T) {
	res, err := testUC.GetPositions("a", 11)
	assert.NoError(t, err)
	assert.Nil(t, res)

	res, err = testUC.GetPositions("TSLA", 1)
	assert.NoError(t, err)
	assert.Nil(t, res)

	res, err = testUC.GetPositions("TSLA", 11)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	for _, p := range res {
		assert.Equal(t, "TSLA", p.Symbol)
		assert.Equal(t, 11, p.HeldBy)
		assert.Equal(t, 210.1, p.PurchasePrice)
		assert.Equal(t, 2.5, p.Qty)
		assert.True(t, time.Now().After(p.PurchasedAt))
	}
}

func TestGetPortfolio(t *testing.T) {
	ps, err := testUC.GetPortfolio(1)
	assert.Nil(t, ps)
	assert.NoError(t, err)

	ps, err = testUC.GetPortfolio(11)
	assert.NotNil(t, ps)
	assert.NoError(t, err)

	for _, p := range ps {
		assert.True(t, (p.Symbol == "TSLA" || p.Symbol == "AAPL"))
		assert.GreaterOrEqual(t, p.PurchasePrice, 210.1)
		assert.GreaterOrEqual(t, p.Qty, 2.5)
		assert.Equal(t, 11, p.HeldBy)
	}
}

func TestGetPortfolioValue(t *testing.T) {
	res, err := testUC.GetPortfolioValue(-1)
	assert.Equal(t, 0.0, res)
	assert.NoError(t, err)

	res, err = testUC.GetPortfolioValue(2)
	assert.NoError(t, err)
	assert.Positive(t, res)
	assert.InDelta(t, 2311.87, res, 2.0)
}
