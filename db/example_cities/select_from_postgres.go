
package main

import (
	"database/sql"
	"fmt"
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
type city struct {
	Country    string `json:"country"`
	City       string `json:"city"`
	Population string `json:"population"`
	Census     string `json:"census"`
}

func (c *city) Print() {
	fmt.Printf("%-20s %-20s %-20s\n", c.City, c.Country, c.Population)
}

func main() {
	creds := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", creds)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT city, country, population, census FROM cities ORDER BY country;")

	for rows.Next() {
		var c city
		err = rows.Scan(&c.City, &c.Country, &c.Population, &c.Census)
		if err != nil {
			log.Fatal(err)
		}
		c.Print()
	}
}
