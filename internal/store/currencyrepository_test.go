package store_test

import (
	"github.com/stretchr/testify/assert"
	"temp/internal/model"
	"temp/internal/store"
	"testing"
)

func TestCurrencyRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("currencies")
	c, err := s.Currency().Create(&model.Currency{
		Name: "RUB",
		Rate: 1000,
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCurrencyRepository_Find(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("currencies")

	name := "aaa"
	_, err := s.Currency().Find(name)
	assert.Error(t, err)

	s.Currency().Create(&model.Currency{
		Name: "aaa",
		Rate: 999,
	})

	c, err := s.Currency().Find(name)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCurrencyRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("currencies")

	newList, err := s.Currency().GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, newList)
}

func TestCurrencyRepository_Delete(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("currencies")

	name := "aaa"
	err := s.Currency().Delete(name)
	assert.Error(t, err)

	s.Currency().Create(&model.Currency{
		Name: "aaa",
		Rate: 999,
	})

	err = s.Currency().Delete(name)
	assert.NoError(t, err)
}

func TestCurrencyRepository_Update(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("currencies")

	name := "bebra"
	rate := 312312.0
	err := s.Currency().Update("aaaa", name, rate)
	assert.Error(t, err)

	s.Currency().Create(&model.Currency{
		Name: "aaaa",
		Rate: 999,
	})

	err = s.Currency().Update("aaaa", name, rate)
	assert.NoError(t, err)
}
