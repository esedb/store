package repository

import (
	"estore/model"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UserRepository encapsulate logic for user
type UserRepository struct {
	C *mgo.Collection
}

//CreateUser creat User in Collection
func (r *UserRepository) CreateUser(user *model.User) error {
	_id := bson.NewObjectId()
	user.Id = _id
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	m := bson.M{
		"_id":       _id,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"email":     user.Email,
		"password":  hpass,
		"roles":     user.Roles,
		"phone":     user.Phone,
	}

	err = r.C.Insert(&m)

	return err

}

//UpdateUser update User in Collection
func (r *UserRepository) UpdateUser(user *model.User) error {
	err := r.C.Update(bson.M{"email": user.Email}, bson.M{
		"$set": bson.M{
			"firstname":  user.FirstName,
			"lastname":   user.LastName,
			"phone":      user.Phone,
			"udpated_at": user.UpdatedAt,
		},
	})

	return err

}

//GetUserByEmail find user by email
func (r *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.C.Find(bson.M{"email": email}).One(&user)

	return user, err
}
