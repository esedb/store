package common

import (
	"gopkg.in/mgo.v2"
)

//Context create context for a particular collection
type Context struct {
	MongoSession *mgo.Session
}

//Close is used to close running session
func (c *Context) Close() {
	c.MongoSession.Close()
}

//DbCollection Get Collection
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(AppConfig.DataBase).C(name)
}

//NewContext create New Context
func NewContext() *Context {
	session := GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}

	return context
}
