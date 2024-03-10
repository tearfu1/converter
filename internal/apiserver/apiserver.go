package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"temp/internal/model"
	"temp/internal/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRouter()

	//if err := s.configureStore(); err != nil {
	//	return err
	//}
	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/getAll", s.handleGetAll()).Methods("GET")
	s.router.HandleFunc("/currencies/{key}", s.handleFind()).Methods("GET")
	s.router.HandleFunc("/create", s.handleCreate()).Methods("POST")
	s.router.HandleFunc("/delete/{key}", s.handleDelete()).Methods("DELETE")
	s.router.HandleFunc("/update/{key}", s.handleUpdate()).Methods("PUT")
	s.router.HandleFunc("/converter", s.handleConverter()).Methods("GET")
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIserver) handleGetAll() http.HandlerFunc {
	currenciesList, err := s.store.Currency().GetAll()
	if err != nil {
		logrus.Fatal(err)
		return nil
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(currenciesList)
	}
}

func (s *APIserver) handleFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		currency, err := s.store.Currency().Find(vars["key"])
		if err != nil {
			logrus.Info(err)
		}
		json.NewEncoder(w).Encode(currency)
	}
}

func (s *APIserver) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var currency model.Currency
		if err := json.NewDecoder(r.Body).Decode(&currency); err != nil {
			logrus.Info(err)
		}
		if _, err := s.store.Currency().Create(&currency); err != nil {
			logrus.Info(err)
		}
		json.NewEncoder(w).Encode(currency)
	}
}

func (s *APIserver) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		if err := s.store.Currency().Delete(vars["key"]); err != nil {
			logrus.Info(err)
		}
		json.NewEncoder(w).Encode(vars["key"])
	}
}

func (s *APIserver) handleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		var currency model.Currency
		if err := json.NewDecoder(r.Body).Decode(&currency); err != nil {
			logrus.Info(err)
		}
		fmt.Println(currency)
		if err := s.store.Currency().Update(vars["key"], currency.Name, currency.Rate); err != nil {
			logrus.Info(err)
		}
		json.NewEncoder(w).Encode(currency)
	}
}

func (s *APIserver) handleConverter() http.HandlerFunc {
	type re struct {
		From   string
		To     string
		Amount float64
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req := new(re)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logrus.Info(err)
		}
		//fmt.Println(req)
		logrus.Info(req.From, req.To, req.Amount)
		modelCurrencyFrom, err := s.store.Currency().Find(req.From)
		if err != nil {
			logrus.Info(err)
		}
		modelCurrencyTo, err := s.store.Currency().Find(req.To)
		if err != nil {
			logrus.Info(err)
		}
		currencyFromNorm := 1 / modelCurrencyFrom.Rate
		currencyToNorm := 1 / modelCurrencyTo.Rate
		amount := currencyFromNorm / currencyToNorm * req.Amount
		logrus.Info(amount)
		json.NewEncoder(w).Encode(fmt.Sprintf("%f", amount))
	}
}
