package common

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var err error

//GetSession Return Sessions
func GetSession() *mgo.Session {
	if session == nil {
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
	}
	return session
}

//CreateSession create mongo session
func createSession() {
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
}

func addIndexes() {
	userIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	session := GetSession().Copy()
	defer session.Close()
	usersCol := session.DB(AppConfig.DataBase).C("users")
	err = usersCol.EnsureIndex(userIndex)

	if err != nil {
		log.Fatalf("[addIndex] - %s\n", err)
	}
}

//StartUp start Dabase configuration Sessions
func StartUp() {
	initConfig()
	createSession()
}
