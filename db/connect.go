package db

import (
    "fmt"
    "gopkg.in/mgo.v2"
    "os"
)

var (
    // Session stores mongo session
    Session *mgo.Session

    // Mongo stores the mongodb connection string information
    Mongo *mgo.DialInfo
)

const (
    MongoDBUrl = "mongodb://:@ds015899.mlab.com:15899/"
)

// Connect connects to mongodb
func Connect() {
    uri := os.Getenv("MONGODB_URL")

    if len(uri) == 0 {
        uri = MongoDBUrl
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
