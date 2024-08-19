package models

// func TestInsertPosition(t *testing.T) {
// 	symbol := "TSLA"
// 	heldBy := 11
// 	qty := 2.5
// 	purchasePrice := 210.1
// 	purchasedAt := time.Now()

// 	id, err := positions.Insert(symbol, heldBy, purchasedAt, purchasePrice, qty)

// 	assert.NoError(t, err)
// 	assert.NotEqual(t, id, -1)
// 	assert.Positive(t, id)
// }

// func TestGetPosition(t *testing.T) {
// 	symbol := "TSLA"
// 	heldBy := 11
// 	qty := 2.5
// 	purchasePrice := 210.1
// 	purchasedAt := time.Now()

// 	id, err := positions.Insert(symbol, heldBy, purchasedAt, purchasePrice, qty)

// 	assert.NoError(t, err)
// 	assert.NotEqual(t, id, -1)
// 	assert.Positive(t, id)

// 	p, err := positions.Get(-1)
// 	assert.EqualError(t, err, sql.ErrNoRows.Error())
// 	assert.Nil(t, p)

// 	p, err = positions.Get(id)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, p)
// 	assert.Equal(t, p.symbol, symbol)
// 	assert.Equal(t, p.heldBy, heldBy)
// 	assert.Equal(t, p.purchasePrice, purchasePrice)
// 	assert.WithinDuration(t, p.purchasedAt, purchasedAt, time.Second)
// 	assert.Equal(t, p.qty, qty)
// }

// func TestGetStocks(t *testing.T) {
// 	stocks, err := positions.GetStock("fake", 11)

// 	assert.Empty(t, stocks)
// 	assert.NoError(t, err)

// 	stocks, err = positions.GetStock("tsla", 0)
// 	assert.Empty(t, stocks)
// 	assert.NoError(t, err)

// 	stocks, err = positions.GetStock("tsla", 11)
// 	assert.NotEmpty(t, stocks)
// 	assert.NoError(t, err)

// 	for _, s := range stocks {
// 		assert.Equal(t, s.symbol, "TSLA")
// 		assert.Equal(t, s.heldBy, 11)
// 		assert.Positive(t, s.ID)
// 	}
// }
