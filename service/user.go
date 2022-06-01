package service

import (
	"echo-app/model"
	"echo-app/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

func (u User) ChangePassword(id primitive.ObjectID, body model.UserChangePassword) (string, error) {
	// check currentPassword
	userBSON, _ := userDAO.GetByID(id)
	if util.CheckPasswordHash(body.CurrentPassword, userBSON.Password) != nil {
		return util.CurrentPasswordIsIncorrect, nil
	}

	// HashPassword before update
	newPassword, _ := util.HashPassword(body.NewPassword)

	// update password
	result, err := userDAO.UpdatePassword(id, newPassword)

	if err != nil || result.ModifiedCount < 1 {
		return util.InvalidData, err
	}

	return util.UpdateSuccessFully, nil
}

func (u User) GetInfo(id primitive.ObjectID) (model.UserInfo, error) {
	var (
		info model.UserInfo
	)

	// get user
	user, err := userDAO.GetInfo(id)
	if err != nil {
		return info, err
	}

	// convert to userInfo
	info = user.ConvertToJSON()

	return info, nil
}

func (u User) UpdateInfo(id primitive.ObjectID, body model.UserUpdate) error {
	bodyBSON := body.ConvertToBSON()
	return userDAO.UpdateInfo(id, bodyBSON)
}
