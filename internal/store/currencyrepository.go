package store

import (
	"fmt"
	"log"
	"temp/internal/model"
)

type CurrencyRepository struct {
	store *Store
}

func (r *CurrencyRepository) Create(c *model.Currency) (*model.Currency, error) {
	err := r.store.db.QueryRow(
		"INSERT INTO `currencies` (`name`, `rate`) VALUES ( ?, ?);", c.Name, c.Rate,
	).Err()
	if err != nil {
		log.Fatalf("impossible insert currency: %s", err)
		return nil, err
	}
	return c, nil
}

func (r *CurrencyRepository) Find(name string) (*model.Currency, error) {
	c := &model.Currency{}
	if err := r.store.db.QueryRow(
		"SELECT `name`, `rate` FROM `currencies` WHERE `name` = ?",
		name,
	).Scan(
		&c.Name,
		&c.Rate,
	); err != nil {
		return nil, err
	}

	return c, nil
}

func (r *CurrencyRepository) GetAll() ([]*model.Currency, error) {
	currenciesList := make([]*model.Currency, 0)

	rows, err := r.store.db.Query("select * from `currencies`;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	currency := &model.Currency{}

	for rows.Next() {
		err := rows.Scan(&currency.Name, &currency.Rate)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		//fmt.Printf("this %s, %f", currency.Name, currency.Rate)
		currenciesList = append(currenciesList, currency)
	}
	return currenciesList, nil
}

func (r *CurrencyRepository) Delete(name string) error {
	_, err := r.Find(name)
	if err != nil {
		return err
	}
	_, err = r.store.db.Query("DELETE FROM `currencies` WHERE `name` = ?", name)
	if err != nil {
		return err
	}
	return nil
}

func (r *CurrencyRepository) Update(src string, name string, rate float64) error {
	_, err := r.Find(src)
	if err != nil {
		return err
	}

	_, err = r.store.db.Query("UPDATE currencies SET `name` = ?, `rate` = ? WHERE `name` = ?;", name, rate, src)
	if err != nil {
		return err
	}
	return nil
}
