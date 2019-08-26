package db

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/mvinturis/mailbox-automation/common/config"
)

var db *mgo.Session

func init() {
	mongoSession, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    config.DB_MONGOADDRS,
		Database: config.DB_NAME,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		fmt.Println("database connection error while connecting to MongoDB: %s", err.Error())
	}

	db = mongoSession
}

func NewSession() *mgo.Database {
	return db.Clone().DB(config.DB_NAME)
}

func RealNewSession() *mgo.Database {
	return db.Copy().DB(config.DB_NAME)
}
