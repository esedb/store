package model

import "gopkg.in/mgo.v2/bson"

//User MerchantUser struct
type User struct {
	ID        bson.ObjectId `json:"_id,omitempty"`
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
	ID          bson.ObjectId     `json:"id,omitempty"`
	Name        string            `json:"name"`
	Price       int64             `json:"price"`
	Description string            `json:"description"`
	Properties  map[string]string `json:"properties"`
	Files       []FileDetails     `json:"files"`
	StoreID     string            `json:"store_id"`
}

//FileDetails file properites
type FileDetails struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Size string `json:"size"`
}

//Store contains relationships with Owner
type Store struct {
	ID         bson.ObjectId `json:"_id,omitempty"`
	Name       string        `json:"name"`
	Location   string        `json:"location"`
	Address    string        `json:"address"`
	Phone      string        `json:"phone"`
	MerchantID string        `json:"merchant_id"`
}

//Category category models
type Category struct {
	ID   bson.ObjectId `json:"_id,omitempty"`
	Name string        `json:"name"`
}
