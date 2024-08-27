package usecase_test

import (
	"testing"

	"github.com/colin-mcl/stocks/pkg/v1/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetQuote(t *testing.T) {
	q, err := testUC.GetQuote("QWODIMQWD")
	assert.Nil(t, q)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrBadSymbol.Error())

	q, err = testUC.GetQuote("AAPL")
	assert.NotNil(t, q)
	assert.NoError(t, err)
	assert.Equal(t, "AAPL", q.Symbol)
	assert.Equal(t, "Apple Inc.", q.ShortName)
	assert.Equal(t, "US", q.Region)
	assert.Equal(t, "EDT", q.TimezoneShort)
	assert.Equal(t, "us_market", q.Market)
	assert.Equal(t, "USD", q.Currency)

}
