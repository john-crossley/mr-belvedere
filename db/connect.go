package db

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "os"
)

var (
    Session *mgo.Session
    Mongo *mgo.DialInfo
)

const (
    MongoDbURI = "mongodb://localhost:27017/mr-belvedere"
)

// Connect connects to mongodb
func Connect() {
    uri := os.Getenv("MONGODB_URL")

    if len(uri) == 0 {
        uri = MongoDbURI
    }

    mongo, err := mgo.ParseURL(uri)
    s, err := mgo.Dial(uri)
    if err != nil {
        fmt.Printf("Can't connect to mongo, go error %v\n", err)
        panic(err.Error())
    }
    s.SetSafe(&mgo.Safe{})
    fmt.Println("Connected to", uri)
    Session = s
    Mongo = mongo
}
