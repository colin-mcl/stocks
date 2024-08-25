package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQuote(t *testing.T) {
	s, err := testClient.GetQuote("")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "invalid quote symbol")
	assert.Equal(t, s, "")
}
