package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
)

//CreateUser createUser service
func CreateUser(user *model.User) (*model.User, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	r := &repo.UserRepository{
		C: c,
	}
	user, err := r.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, err

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
func UpdateUser(user *model.User) (*model.User, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	r := &repo.UserRepository{
		C: c,
	}
	user, err := r.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, err

}
