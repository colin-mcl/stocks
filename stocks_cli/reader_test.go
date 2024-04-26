package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyTicker(t *testing.T) {
	res, err := handleGetRequest("")

	require.Nil(t, res)
	require.NotEmpty(t, err)
	require.Error(t, err)
}

func TestInvalidTicker(t *testing.T) {
	res, err := handleGetRequest("ZZZZZZZZ")

	require.Nil(t, res)
	require.NotEmpty(t, err)
	require.Error(t, err)
	require.Equal(t, "Ticker ZZZZZZZZ not found.\n", err.Error())

}

func TestValidTicker(t *testing.T) {
	res, err := handleGetRequest("AAPL")

	require.Nil(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, "Apple Inc.", res.ShortName)
	require.Equal(t, "AAPL", res.Symbol)
	require.Equal(t, "US", res.Region)
}
