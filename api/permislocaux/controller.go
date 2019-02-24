package permislocaux

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

func HandlePermislocauxGetAll(writer http.ResponseWriter, request *http.Request) {
	var results []*Permislocaux
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
		var elem Permislocaux
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
	writer.Header().Set("Context-Type", "application/json")
	writer.Write(output)
	//writer.Write([]byte("permislocaux"))
}

func HandlePermislocauxGetById(writer http.ResponseWriter, request *http.Request) {
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

	writer.Write([]byte("eee"))
}

