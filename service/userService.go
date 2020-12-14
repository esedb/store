package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
	"time"
)

//CreateUser createUser service
func CreateUser(user *model.User) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	r := &repo.UserRepository{
		C: c,
	}
	user.CreatedAt = time.Now().UTC().Unix()

	err := r.CreateUser(user)

	return err

}

//GetUserByEmail getUser by email
func GetUserByEmail(email string) (*model.User, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	r := &repo.UserRepository{
		C: c,
	}
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return &user, err

}

//UpdateUser update user by email
func UpdateUser(user *model.User) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	r := &repo.UserRepository{
		C: c,
	}
	user.UpdatedAt = time.Now().Unix()
	err := r.UpdateUser(user)

	return err
}
