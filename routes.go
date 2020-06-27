package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//=======Functions used in the routs ==============

func listAllPlanets(w http.ResponseWriter, r *http.Request) {
	var planets []Planet

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := planetsColl.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &planets)

	jPlanets, _ := json.Marshal(planets)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	fmt.Fprintf(w, "%s\n", jPlanets)
}

func findPlanet(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	var planet Planet
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if queries["id"] != nil {
		id := queries["id"][0]
		objID, _ := primitive.ObjectIDFromHex(id)

		planetsColl.FindOne(ctx, bson.M{"_id": objID}).Decode(&planet)

	} else if queries["nome"] != nil {
		name := queries["nome"][0]
		planetsColl.FindOne(ctx, bson.M{"nome": name}).Decode(&planet)

	} else {
		fmt.Fprintf(w, "The query is not valid\n")
		return
	}

	// if planet.Nome == "" {
	// 	fmt.Fprintf(w, "The querie found no results \n")
	// 	return
	// }

	jPlanet, _ := json.Marshal(planet)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200

	fmt.Fprintf(w, "%s\n", jPlanet)
}

func insertPlanet(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var planet Planet

	if !json.Valid(body) {
		fmt.Fprintf(w, "Data is not valid JSON\n")
		return
	}
	json.Unmarshal(body, &planet)

	//====Inserting Stuff=======
	var err error

	planet.ID = primitive.NewObjectID()
	planet.Filmes, err = getFilms(planet.Nome)

	if err != nil {
		fmt.Fprintf(w, "The named planet does not exist in the sw universe \n")
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	planetResult, err := planetsColl.InsertOne(ctx, planet)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "inserted ID = %v, data=%v \n", planetResult.InsertedID, planet)
}

func deletePlanet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	objID, err := primitive.ObjectIDFromHex(query["id"][0])
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err != nil {
		log.Println(err)
	}
	results, err := planetsColl.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Deleted %v from the database\n", results.DeletedCount)
}
