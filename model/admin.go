package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Admin struct {
		ID             primitive.ObjectID `bson:"_id" `
		Email          string             `bson:"email"`
		Username       string             `bson:"username"`
		HashedPassword string             `bson:"hashedPassword"`
		FullName       string             `bson:"fullName"`
		DateOfBirth    string             `bson:"dateOfBirth"`
		Avatar         string             `bson:"avatar"`
		Gender         string             `bson:"gender"`
		Phone          string             `bson:"phone"`
		CreatedAt      string             `bson:"createdAt"`
		UpdatedAt      string             `bson:"updatedAt"`
	}
	AdminResponse struct {
		ID             primitive.ObjectID `json:"_id"`
		Email          string             `json:"email"`
		Username       string             `json:"username"`
		Hashedpassword string             `json:"hashedpassword"`
		Fullname       string             `json:"fullname"`
		DateOfBirth    string             `json:"dateOfBirth"`
		Avatar         string             `json:"avatar"`
		Gender         string             `json:"gender"`
		Phone          string             `json:"phone"`
		CreatedAt      string             `json:"createdAt"`
		UpdatedAt      string             `json:"updatedAt"`
	}

	AdminLoginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (a AdminLoginBody) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Username, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
}
