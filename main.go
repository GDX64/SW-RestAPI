package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const apiURI = "mongodb://localhost:27017"
const port = 3000

var ctx context.Context
var client *mongo.Client
var planetsColl *mongo.Collection

func init() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(apiURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	test := client.Database("test")
	planetsColl = test.Collection("planets")
}

func main() {
	defer client.Disconnect(ctx)

	r := mux.NewRouter()

	r.HandleFunc("/planet", insertPlanet).Methods("POST")
	r.HandleFunc("/planets", listAllPlanets).Methods("GET")
	r.HandleFunc("/planet/", findPlanet).Methods("GET")
	r.HandleFunc("/planet/", deletePlanet).Methods("DELETE")

	log.Printf("Server starting at %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
