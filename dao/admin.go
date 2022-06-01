package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct{}

// FindByUsername
func (Admin) FindByUsername(username string) (model.Admin, error) {
	var (
		admin model.Admin
	)

	// find
	filter := bson.M{
		"username": username,
	}

	if err := database.AdminCol().FindOne(util.Ctx, filter).Decode(&admin); err != nil {
		return admin, err
	}

	return admin, nil
}

func (Admin) InitAdminAccount() {
	count, _ := database.AdminCol().CountDocuments(util.Ctx, bson.D{})

	if count == 0 {
		admin := model.Admin{
			ID:             primitive.NewObjectID(),
			Email:          "admin123@gmail.com",
			Username:       "admin123",
			HashedPassword: "123456",
			FullName:       "admin",
			DateOfBirth:    "2022-05-10",
			Gender:         "Male",
			Phone:          "0779925153",
		}
		database.AdminCol().InsertOne(util.Ctx, admin)
	}
}

func (Admin) ProfileFindByID(id primitive.ObjectID) (model.Admin, error) {
	var (
		adminProfile model.Admin
	)

	// find profile
	filter := bson.M{"_id": id}

	if err := database.AdminCol().FindOne(util.Ctx, filter).Decode(&adminProfile); err != nil {
		return adminProfile, err
	}

	return adminProfile, nil
}

func (Admin) UpdateMyProfile(id primitive.ObjectID, newProfile model.Admin) error {
	update := model.Admin{
		FullName:    newProfile.FullName,
		DateOfBirth: newProfile.DateOfBirth,
		Gender:      newProfile.Gender,
		Phone:       newProfile.Phone,
	}

	// UpdateOne
	if _, err := database.AdminCol().UpdateOne(util.Ctx, bson.M{"_id": id}, bson.M{"$set": update}); err != nil {
		return err
	}
	return nil
}
