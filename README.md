### dependencies use : go mod

- from GO >=1.11 we can develop outside $GOPATH
- run `go mod init {module_name}` will create **_go.mod_** file
- run `go get {package}` will update **_go.mod_**

### impact on local importations

- when importing local packages as ./mypackage/mysubpackage

`import "{module_declared_in_gomod}/mypackage/mysubpackage"`

- do not forget to prefix with the name of the main module 

### mongo driver

- replace mgo with mongo-driver-go from mongodb 
- use mongo-driver-go/bson/primitive to handle mongo types especially objectID
- use mongo-driver-go/bson tags to match mongo fields

### router

- replace gorilla with go-chi
- go-chi comes with a bunch of middlewares (logger, jwt, requestID...)

### hot reloading

### configure the project with viper