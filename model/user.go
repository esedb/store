package model

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Roles     string `json:"roles"`
}

type Product struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}
