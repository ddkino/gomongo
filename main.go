package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/go-chi/chi"
)

type Permis struct {
	Siret                  string `json:"siret"`
	Codepostaledudemandeur string `json:"codepostal"`
}

type todo struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Task       string    `json:"task"`
}

//noinspection ALL
func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	r.Post("/permislocaux", func(w http.ResponseWriter, r *http.Request) {
		client, err := mongo.NewClient("mongodb://localhost:27017/kb")
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		collection := client.Database("kb").Collection("test")
		res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		id := res.InsertedID
		fmt.Println(id)
		w.Write([]byte(("rrrr")))
	})

	r.Get("/permislocaux", func(w http.ResponseWriter, r *http.Request) {
		client, err := mongo.NewClient("mongodb://localhost:27017/kb")
		if err != nil {
			fmt.Println(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		collection := client.Database("kb").Collection("permislocaux")
		filter := bson.D{{"region", "84"}, {"anneedepot", "2016"}}
		cur, err := collection.Find(ctx, filter)

		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
		}
		w.Write([]byte("permislocaux"))
	})

	r.Get("/permislocaux/id", func(w http.ResponseWriter, r *http.Request) {
		client, err := mongo.NewClient("mongodb://localhost:27017/kb")
		if err != nil {
			fmt.Println(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		collection := client.Database("kb").Collection("permislocaux")
		var result struct {
			Siret string
		}
		filter := bson.M{}
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		out, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		w.Write([]byte("eee"))
	})

	r.Get("/permislocaux/json", func(w http.ResponseWriter, r *http.Request) {
		var results []*Permis
		client, err := mongo.NewClient("mongodb://localhost:27017/kb")
		if err != nil {
			fmt.Println(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		collection := client.Database("kb").Collection("permislocaux")
		filter := bson.D{{"region", "84"}, {"anneedepot", "2016"}}
		cur, err := collection.Find(ctx, filter)

		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var elem Permis
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, &elem)
			//fmt.Println(elem)
		}
		output, err := json.Marshal(results)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Context-Type", "application/json")
		w.Write(output)
	})

	r.Get("/test/id", func(w http.ResponseWriter, r *http.Request) {
		client, err := mongo.NewClient("mongodb://localhost:27017/kb")
		if err != nil {
			fmt.Println(err)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		collection := client.Database("kb").Collection("test")
		var result struct {
			Value float64
		}
		filter := bson.M{"name": "pi"}
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		out, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		w.Write([]byte("eee"))
	})

	type Profile struct {
		Name    string
		Hobbies []string
	}
	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {

		profile := Profile{
			Name:    "dede",
			Hobbies: []string{"aaa", "ppp"},
		}

		output, err := json.Marshal(profile)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Context-Type", "application/json")
		w.Write(output)
	})

	http.ListenAndServe(":3333", r)

}
