package conf

import (
	"flag"
	"log"

	"gopkg.in/mgo.v2"
)

var (
	addr        = flag.String("addr", ":3000", "Server Address")
	mongoAddr   = flag.String("mongo", "127.0.0.1:27017", "Mongodb Address")
	dbName      = flag.String("db", "stack", "Database Name")
	tokenExpire = flag.Int64("expire-token", 3, "Expire User Access Token After Hours")
)

var (
	Addr string
	M    *mgo.Session
	MDB  *mgo.Database
)

func Load() {
	// advanced logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// parse commandline flags
	flag.Parse()

	// init configs for global access
	Addr = *addr
	M = dialMongo()
	MDB = getMDB()
}
