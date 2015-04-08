package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"log"
	"os"
)

var blacklist = []string{"favicon.ico", "apple-touch-icon.png"}

type Search struct {
	Search string
}

func initDB() (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		panic(err)
	}

	return session.DB("heroku_app35656076").C("searches"), session
}

func SaveSearch(search string) {
	collection, session := initDB()

	defer session.Close()

	err := collection.Insert(&Search{search})
	if err != nil {
		log.Println(err)
	}
}

func TopSearches() string {
	collection, session := initDB()

	defer session.Close()

	pipe := collection.Pipe(
		[]bson.M{
			{
				"$group": bson.M{
					"_id":   bson.M{"search": "$search"},
					"total": bson.M{"$sum": 1},
				},
			},
			{
				"$sort": bson.M{"total": -1},
			},
		})

	results := []bson.M{}
	result := ""

	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}

	for _, res := range results {
		search := res["_id"].(bson.M)
		if !stringInSlice(search["search"].(string), blacklist) {
			result += fmt.Sprintf("Search: <a href=\"/%s\">%s</a>   %d</br>", search["search"], search["search"], res["total"])
		}
	}

	return result
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
