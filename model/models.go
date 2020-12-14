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
	CreatedAt int64         `json:"created_at"`
	UpdatedAt int64         `json:"updated_at"`
}

//Product Used for create Products
//properties field is other generic properties from the main one
type Product struct {
	Id          bson.ObjectId     `bson:"_id" json:"id"`
	Name        string            `json:"name"`
	Price       float64           `json:"price"`
	Description string            `json:"description"`
	Properties  map[string]string `json:"properties"`
	Files       []FileDetails     `json:"files"`
	StoreID     string            `json:"store_id"`
	CreatedAt   int64             `json:"created_at"`
	UpdatedAt   int64             `json:"updated_at"`
}

//FileDetails file properites
type FileDetails struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Size string `json:"size"`
}

//Store contains relationships with Owner
type Store struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `json:"name"`
	Location   string        `json:"location"`
	Address    string        `json:"address"`
	Phone      string        `json:"phone"`
	MerchantID string        `json:"merchant_id"`
}

//Category category models
type Category struct {
	Id   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `json:"name"`
}
