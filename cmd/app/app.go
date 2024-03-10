package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:"-"`
	Gesmes  string   `xml:"-"`
	Xmlns   string   `xml:"-"`
	Subject string   `xml:"-"`
	Sender  struct {
		Text string `xml:"-"`
		Name string `xml:"-"`
	} `xml:"-"`
	Cube struct {
		Text string `xml:"-"`
		Cube struct {
			Text string `xml:"-"`
			Time string `xml:"-"`
			Cube []struct {
				Text     string  `xml:"-"`
				Currency string  `xml:"currency,attr"`
				Rate     float64 `xml:"rate,attr"`
			} `xml:"Cube"`
		} `xml:"Cube"`
	} `xml:"Cube"`
}

func (e Envelope) String() string {
	return fmt.Sprintf("Envelope is %v",
		e.Cube)
}

var (
	dsn string
)

func init() {
	flag.StringVar(&dsn, "dsn", "example_user:example_user_password@tcp(mysql:3306)/example_db", "data source name")
	//example_user:example_user_password@tcp(localhost:3306)/example_db
	//	user=username password=password dbname=example_db host=localhost port=5432 sslmode=disable
}

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	db.Query("CREATE TABLE IF NOT EXISTS `currencies` ( `name` varchar(10), `rate` float);")

	if err := db.QueryRow(
		"INSERT INTO `currencies` (`name`, `rate`) VALUES (?, ?);", "EUR", 1.0,
	).Err(); err != nil {
		log.Fatalf("impossible insert currency: %s", err)
		return
	}

	resp, _ := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	bytes, _ := io.ReadAll(resp.Body)
	string_body := string(bytes)
	resp.Body.Close()

	envlp := new(Envelope)
	if err := xml.Unmarshal([]byte(string_body), envlp); err != nil {
		panic(err)
	}
	currencies := envlp.Cube.Cube.Cube
	for _, currency := range currencies {
		currency_name := currency.Currency
		currency_rate := currency.Rate

		query := "INSERT INTO `currencies` (`name`, `rate`) VALUES ( ?, ?)" // VALUES($1, $2)
		_, err := db.ExecContext(context.Background(), query, currency_name, currency_rate)
		if err != nil {
			log.Fatalf("impossible insert test: %s", err)
		}
	}

	fmt.Println("Table was filled")
	//db.Query("DROP TABLE `test`;")
}
