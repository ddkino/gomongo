package programmesneufs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	optionsMongo "go.mongodb.org/mongo-driver/mongo/options"
)

const (
	db         = "kb"
	collection = "programmesseloger_copy"
)

func HandleProgrammesneufsTest(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("programmesseloger"))
}

/*
HandleProgrammesneufsGetAll retrieves all datas by date (from)
*/
func HandleProgrammesneufsFindByDate(writer http.ResponseWriter, request *http.Request) {
	//writer.Write([]byte("programmesseloger"))
	var results []*Programmesneufs
	client, err := mongo.NewClient(optionsMongo.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// Read body
	b, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, "ReadAll="+err.Error(), 500)
		return
	}

	// Unmarshal
	var msg map[string]interface{}
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(writer, "Unmarshal="+err.Error(), 500)
		return
	}

	var tMin, tMax time.Time
	var first int64
	var lon, lat float64

	if _, ok := msg["lon"]; ok {
		lon, err = strconv.ParseFloat(msg["lon"].(string), 64)
		if err != nil {
			lon = 5
		}
	} else {
		lon = 5
	}

	if _, ok := msg["lat"]; ok {
		lat, err = strconv.ParseFloat(msg["lat"].(string), 64)
		if err != nil {
			lat = 45
		}
	} else {
		lat = 45
	}

	if _, ok := msg["first"]; ok {
		first, err = strconv.ParseInt(msg["first"].(string), 10, 32)
		if err != nil {
			first = 100
		}
	} else {
		first = 100
	}

	if _, ok := msg["tMin"]; ok {
		tMin, err = time.Parse("2006-01-02", msg["tMin"].(string))
		if err != nil {
			tMin, _ = time.Parse("2006-01-02", "2018-01-01")
		}
	} else {
		tMin, _ = time.Parse("2006-01-02", "2018-01-01")
	}

	if _, ok := msg["tMax"]; ok {
		tMax, err = time.Parse("2006-01-02", msg["tMax"].(string))
		if err != nil {
			tMax, _ = time.Parse("2006-01-02", "2222-01-01")
		}
	} else {
		tMax, _ = time.Parse("2006-01-02", "2222-01-01")
	}

	fmt.Println(msg)
	fmt.Println(tMax)
	options := optionsMongo.Find()
	options.SetLimit(first)

	collection := client.Database(db).Collection(collection)
	filter := bson.M{
		"created_at": bson.M{"$gte": tMin, "$lte": tMax},
		"geojson": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lon, lat},
				},
				"$maxDistance": 100000,
			},
		},
	}
	cur, err := collection.Find(ctx, filter, options)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem Programmesneufs
		// var elem interface{}
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
	writer.Header().Set("Context-Type", "application/json; charset=utf-8")
	writer.Write(output)
}

/*
HandleProgrammesneufsGetAll retrieves all datas by date (from)
*/
func HandleProgrammesneufsGetAll(writer http.ResponseWriter, request *http.Request) {
	var results []*Programmesneufs
	client, err := mongo.NewClient(optionsMongo.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// var dat map[string]interface{}
	// json.Unmarshal(request.Body, &dat)
	fmt.Println(request.Body)

	options := optionsMongo.Find()
	options.SetLimit(1000)

	collection := client.Database(db).Collection(collection)
	filter := bson.M{}
	cur, err := collection.Find(ctx, filter, options)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem Programmesneufs
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
	writer.Header().Set("Context-Type", "application/json; charset=utf-8")
	writer.Write(output)
	//writer.Write([]byte("permislocaux"))
}

/*
db.programmesseloger_copy.aggregate([
	{ $addFields: {
			"created_at": {
					"$toDate": "$creationDate"
			}
			}
	},
	{ $addFields: {
			"geo": { type: "Point", coordinates: [ 40, 5 ] }
			}
	},
	{ $match: {created_at: {$gte: new Date("2018-10-01")}}},
	{ $match: {created_at: {$lte: new Date("2019-01-01")}}},
	$nearSphere: {$geometry: { type: "Point", coordinates: [ 40, 5 ]}}, $minDistance: 1, $maxDistance: 10000000 },
	{ $limit: 10 },
]
)
*/

// func HandlePermislocauxGetById(writer http.ResponseWriter, request *http.Request) {
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)

// 	collection := client.Database("kb").Collection("permislocaux")
// 	var result struct {
// 		Siret string
// 	}
// 	filter := bson.M{}
// 	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
// 	err = collection.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	out, err := json.Marshal(result)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(out))

// 	writer.Write([]byte("eee"))
// }
