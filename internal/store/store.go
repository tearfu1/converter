package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
)

type Store struct {
	config             *Config
	db                 *sql.DB
	CurrencyRepository *CurrencyRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Currency() *CurrencyRepository {
	if s.CurrencyRepository != nil {
		return s.CurrencyRepository
	}
	s.CurrencyRepository = &CurrencyRepository{
		store: s,
	}
	return s.CurrencyRepository
}
