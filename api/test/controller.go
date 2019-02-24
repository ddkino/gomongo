package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func handleProfileGetJson(writer http.ResponseWriter, request *http.Request) {
	profile := Profile{
		Name:    "dede",
		Hobbies: []string{"aaa", "ppp"},
	}

	output, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}
	writer.Header().Set("Context-Type", "application/json")
	writer.Write(output)
}

func handleProfileGetByName(writer http.ResponseWriter, request *http.Request) {
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
	output, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Context-Type", "application/json")
	writer.Write(output)
}

func handleProfileInsertOne(writer http.ResponseWriter, request *http.Request) {
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
	writer.Write([]byte(("rrrr")))
}
