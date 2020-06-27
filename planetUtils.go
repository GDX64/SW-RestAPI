package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Planet struct {
	ID      primitive.ObjectID `bson:"_id",omitempty`
	Nome    string             `bson:"nome",omitempty`
	Clima   string             `bson:"clima",omitempty`
	Terreno string             `bson:"terreno",omitempty`
	Films   int                `bson:"films",omitempty`
}

func getFilms(name string) (movies int, err error) {
	resp, err := http.Get(fmt.Sprintf("https://swapi.dev/api/planets/?search=%v", name))

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from getFilms")
			err = errors.New("No movie with this planet")
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if json.Valid(body) {
		var doc map[string]interface{}
		json.Unmarshal(body, &doc)
		films := doc["results"].([]interface{})[0].(map[string]interface{})["films"].([]interface{})

		movies = len(films)
	}
	return movies, err
}

func getPodcast(collection *mongo.Collection, query bson.M) []Planet {
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		panic(err)
	}
	var pd []Planet
	cursor.All(ctx, &pd)

	return pd
}
