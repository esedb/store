package common

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var err error

//GetSession Return Sessions
func GetSession() *mgo.Session {
	host := os.Getenv("MONGO_HOST")
	address := ""
	if host != "" {
		address = host + ":27017"
	} else {
		address = "localhost:27017"

	}
	if session == nil {
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs: []string{address},

			// Username: AppConfig.DBUser,
			// Password: AppConfig.DBPwd,
			Timeout: 60 * time.Second,
		})
	}
	return session
}

//CreateSession create mongo session
func createSession() {
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{AppConfig.MongoDBHost},
		// Username: AppConfig.DBUser,
		// Password: AppConfig.DBPwd,
		Timeout: 60 * time.Second,
	})

}

func addIndexes() {
	userIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	storeIndex := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	session := GetSession().Copy()
	defer session.Close()
	usersCol := session.DB(AppConfig.DataBase).C("users")
	storesCol := session.DB(AppConfig.DataBase).C("stores")
	err = usersCol.EnsureIndex(userIndex)
	err = storesCol.EnsureIndex(storeIndex)

	if err != nil {
		log.Fatalf("[addIndex] - %s\n", err)
	}
}

//StartUp start Dabase configuration Sessions
func StartUp() {
	initConfig()
	createSession()
	addIndexes()
}
