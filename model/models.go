package model

import "gopkg.in/mgo.v2/bson"

//User MerchantUser struct
type User struct {
	Id        bson.ObjectId `json:"_id,omitempty"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Roles     string        `json:"roles"`
	Phone     string        `json:"phone"`
}

//Product Used for create Products
//properties field is other generic properties from the main one
type Product struct {
	Name        string            `json:"name"`
	Price       int64             `json:"price"`
	Description string            `json:"description"`
	Properties  map[string]string `json:"properties"`
}
