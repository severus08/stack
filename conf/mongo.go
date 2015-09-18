package conf

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

func dialMongo() *mgo.Session {
	mongo, err := mgo.Dial(*mongoAddr)
	if err != nil {
		log.Fatalln(err)
	}
	return mongo
}

func getMDB() *mgo.Database {
	db := M.DB(*dbName)

	uc := db.C("tokens")
	if err := uc.EnsureIndex(mgo.Index{
		Key:         []string{"created_at"},
		ExpireAfter: time.Hour * time.Duration(*tokenExpire),
	}); err != nil {
		log.Fatalln(err)
	}

	return db
}
