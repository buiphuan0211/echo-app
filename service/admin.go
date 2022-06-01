package service

import (
	"echo-app/model"
	"echo-app/util"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct{}

func (a Admin) AdminLogin(body model.AdminLoginBody) (string, error) {
	// find admin in db
	admin, err := adminDAO.FindByUsername(body.Username)
	if err != nil {
		return "", err
	}

	// verify admin password
	if admin.HashedPassword != body.Password {
		return "", errors.New("wrong password")
	}

	data := map[string]interface{}{
		"id":      admin.ID,
		"isAdmin": true,
	}

	// return JWT token
	token, err := util.GenerateToken(data)
	if err != nil {
		return "", err
	}
	return token, err
}

func (a Admin) GetMyProfile(id primitive.ObjectID) (model.Admin, error) {
	return adminDAO.ProfileFindByID(id)
}

func (a Admin) UpdateMyProfile(id primitive.ObjectID, newProfile model.Admin) error {
	return adminDAO.UpdateMyProfile(id, newProfile)
}
