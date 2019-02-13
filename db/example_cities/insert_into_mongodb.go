
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"io/ioutil"
	"log"
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

	// mongo part

	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("mondial").Collection("cities")
	for i, c := range cities {
		fmt.Println(i, c.City, c.Country)
		_, err := collection.InsertOne(context.TODO(), c)
		if err != nil {
			log.Fatal(err)
		}
	}
}
