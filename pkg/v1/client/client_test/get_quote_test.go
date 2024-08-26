package client_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
)

func TestGetQuote(t *testing.T) {
	s, err := testClient.GetQuote("")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid quote symbol")
	assert.Equal(t, s, "")

	r := strings.ToUpper(util.RandomString(20))
	s, err = testClient.GetQuote(r)
	assert.Error(t, err)
	assert.ErrorContains(t, err, fmt.Sprintf("invalid quote symbol: %s", r))
	assert.Equal(t, s, "")

	s, err = testClient.GetQuote("tsla")
	assert.NoError(t, err)
	assert.NotEqual(t, s, "")
	assert.Contains(t, s, "TSLA")
	assert.Contains(t, s, "Tesla")
	assert.Contains(t, s, "USD")

	s, err = testClient.GetQuote("AAAAAAAAA")
	assert.Error(t, err)
	assert.Equal(t, "", s)
}
