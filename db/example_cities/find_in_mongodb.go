
package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

func main() {

	type city struct {
		Country    string `json:"country"`
		City       string `json:"city"`
		Population string `json:"population"`
		Census     string `json:"census"`
	}

	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("mondial").Collection("cities")

	// No filter. There must be a filter, even if it is empty. nil does
	// not work.
	filter := struct{}{}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var c city
		err = cur.Decode(&c)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("| %-20s | %-20s | %s |\n", c.City, c.Country, c.Population)
	}
}
