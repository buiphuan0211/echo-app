package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

func (u User) Register(doc model.UserBSON) (primitive.ObjectID, error) {
	// Insert one
	result, err := database.UserCol().InsertOne(util.Ctx, doc)
	return result.InsertedID.(primitive.ObjectID), err
}

func (u User) GetByUsername(username string) (model.UserBSON, error) {
	var (
		user model.UserBSON
	)

	// filter
	filter := bson.M{"username": username}

	// FindOne
	if err := database.UserCol().FindOne(util.Ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (u User) GetByID(id primitive.ObjectID) (model.UserBSON, error) {
	var (
		user model.UserBSON
	)

	err := database.UserCol().FindOne(util.Ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u User) UpdatePassword(id primitive.ObjectID, newPassword string) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{{"password", newPassword}}},
	}

	return database.UserCol().UpdateOne(util.Ctx, filter, update)
}

func (u User) GetInfo(id primitive.ObjectID) (model.UserBSON, error) {
	var (
		user model.UserBSON
	)

	filter := bson.M{"_id": id}

	if err := database.UserCol().FindOne(util.Ctx, filter).Decode(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (u User) UpdateInfo(id primitive.ObjectID, body model.UserInfoBSON) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": body}

	_, err := database.UserCol().UpdateOne(util.Ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
