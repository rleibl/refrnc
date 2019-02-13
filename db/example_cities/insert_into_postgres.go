
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "test"
  password = "password"
  dbname   = "test"
)

func main() {
	b, err := ioutil.ReadFile("cities.json")
	if err != nil {
		log.Fatal(err)
	}

	type city struct {
		Country    string `json:"country"`
		City       string `json:"city"`
		Population string `json:"population"`
		Census     string `json:"census"`
	}

	var cities []city
	err = json.Unmarshal(b, &cities)
	if err != nil {
		log.Fatal(err)
	}

	// postgres part

	creds := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", creds)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, c := range cities {
		sql := "INSERT INTO cities (city, country, population, census) VALUES ($1, $2, $3, $4)"
		_, err := db.Exec(sql, c.City, c.Country, c.Population, c.Census)
		if err != nil {
			log.Fatal(err)
		}
	}
}
