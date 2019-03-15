### dependencies use : go mod

- from GO >=1.11 we can develop outside $GOPATH
- after upgrade bin go run `cd $GOPATH/src && go get -u -v ./...`
- run `go mod init {module_name}` will create **_go.mod_** file
- run `go get {package}` will update **_go.mod_**

### impact on local importations

- when importing local packages as ./mypackage/mysubpackage

`import "{module_declared_in_gomod}/mypackage/mysubpackage"`

- do not forget to prefix with the name of the main module 

### mongo driver

- replace mgo with mongo-driver-go from mongodb 
- mgo is unmaintained since 07/2018
- use mongo-driver-go/bson/primitive to handle mongo types especially objectID
- use mongo-driver-go/bson tags to match mongo fields
- do not put "@version" `cd $GOPATH`
- download: `go get go.mongodb.org/mongo-driver/mongo`
- install:  `go install go.mongodb.org/mongo-driver/mongo`

### router

- replace gorilla with go-chi
- go-chi comes with a bunch of middlewares (logger, jwt, requestID...)

### hot reloading

### configure the project with viper


### mongodb reformat data
- create index created_at
`db.getCollection('programmesseloger').createIndex({"created_at": 1})`
- create index geoJson
`db.getCollection('programmesseloger').createIndex({"geojson": "2dsphere"})`

- reformat creationDate(string) to created_at(date) and coordinates

```
db.getCollection('programmesseloger').find({}).limit().forEach(
function(item){
    item.created_at = new Date(item.creationDate)
    newFields = {}
    newFields["created_at"]=new Date(item.creationDate);
    if(item.coordinates && item.coordinates.lat['$numberDouble'] && item.coordinates.lon['$numberDouble']) {
        lat = item.coordinates.lat['$numberDouble'];
        lon = item.coordinates.lon['$numberDouble'];
        print(lon, lat)
        newFields["geojson"]={
                type: 'Point',
                coordinates: [lon, lat] // lon , lat
        }     
    } else {
        newFields["geojson"]=null 
    }
    db.getCollection('programmesseloger').update(
        { _id: item._id }, 
        { "$set": newFields }
        )
}
)
```

- check geo query

```
db.getCollection('programmesseloger_copy').find({
    geojson: {
        $near: {
            $geometry: {type: "Point", coordinates: [5.48642875532712, 45.849067]},
            $maxDistance: 100000,
            }
        }
    }).limit(1000)
```
### mongo import CSV 
`mongoimport --db kb --collection open_siret_mars2019 --type csv StockEtablissement_utf8.csv`