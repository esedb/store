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
type ananymous interface{}
type Product struct {
	Id          bson.ObjectId          `bson:"_id" json:"id"`
	Name        string                 `json:"name" validate:"required"`
	Price       float64                `json:"price" validate:"required"`
	Description string                 `json:"description" validate:"required"`
	Properties  []map[string]ananymous `json:"properties"`
	Files       []FileDetails          `json:"files"`
	StoreID     string                 `bson:"store_id" json:"store_id" validate:"required"`
	CreatedAt   int64                  `bson:"created_at" json:"created_at"`
	UpdatedAt   int64                  `bson:"updated_at" json:"updated_at"`
	Qty         int64                  `json:"qty" validate:"required"`
}

//FileDetails file properites
type FileDetails struct {
	Name string  `json:"name" validate:"required"`
	Data string  `json:"data" validate:"required"`
	Size float32 `json:"size" validate:"required"`
}

//Store contains relationships with Owner
type Store struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `json:"name"`
	Location   string        `json:"location"`
	Address    string        `json:"address"`
	Phone      string        `json:"phone"`
	MerchantId string        `bson:"merchant_id" json:"merchant_id"`
}

//Category category models
type Category struct {
	Id   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `json:"name" validate:"required"`
}
