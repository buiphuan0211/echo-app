package model

import (
	"echo-app/util"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserBSON struct {
		ID          primitive.ObjectID `bson:"_id"`
		Email       string             `bson:"email" `
		Username    string             `bson:"username"`
		Password    string             `bson:"password"`
		FullName    string             `bson:"fullName"`
		DateOfBirth string             `bson:"dateOfBirth"`
		Avatar      string             `bson:"avatar"`
		Gender      string             `bson:"gender"`
		Phone       string             `bson:"phone"`
		Address     string             `bson:"address"`
		Status      string             `bson:"status"`
		CreatedAt   time.Time          `bson:"createdAt"`
		UpdatedAt   time.Time          `bson:"updatedAt"`
	}

	UserInfoBSON struct {
		FullName    string `bson:"fullName"`
		Email       string `bson:"email" `
		DateOfBirth string `bson:"dateOfBirth"`
		Gender      string `bson:"gender"`
		Phone       string `bson:"phone"`
		Address     string `bson:"address"`
		Status      string `bson:"status"`
	}

	UserResponse struct {
		ID          primitive.ObjectID `json:"_id"`
		Email       string             `json:"email"`
		Username    string             `json:"username"`
		Password    string             `json:"password"`
		FullName    string             `json:"fullName"`
		DateOfBirth string             `json:"dateOfBirth"`
		Avatar      string             `json:"avatar"`
		Gender      string             `json:"gender"`
		Phone       string             `json:"phone"`
		Address     string             `bson:"address"`
		Status      string             `json:"status"`
		CreatedAt   time.Time          `json:"createdAt"`
		UpdatedAt   time.Time          `json:"updatedAt"`
	}

	UserLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserChangePassword struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
	}

	UserUpdate struct {
		FullName    string `json:"fullName"`
		Email       string `json:"email"`
		Phone       string `json:"phone"`
		DateOfBirth string `json:"dateOfBirth"`
		Gender      string `json:"gender"`
		Address     string `json:"address"`
	}

	UserRegister struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		FullName    string `json:"fullName"`
		Gender      string `json:"gender"`
		DateOfBirth string `json:"dateOfBirth"`
		Phone       string `json:"phone"`
		Address     string `json:"address"`
		Status      string `json:"status"`
	}

	UserInfo struct {
		ID          primitive.ObjectID `json:"_id"`
		FullName    string             `json:"fullName"`
		Email       string             `json:"email"`
		Username    string             `json:"username"`
		DateOfBirth string             `json:"dateOfBirth"`
		Avatar      string             `json:"avatar"`
		Gender      string             `json:"gender"`
		Phone       string             `json:"phone"`
		Address     string             `json:"address"`
	}
)

// ConvertToBSON
func (u UserRegister) ConvertToBSON() UserBSON {
	result := UserBSON{
		ID:          primitive.NewObjectID(),
		Username:    u.Username,
		Password:    u.Password,
		Email:       u.Email,
		FullName:    u.FullName,
		Gender:      u.Gender,
		DateOfBirth: u.DateOfBirth,
		Phone:       u.Phone,
		Address:     u.Address,
		Status:      util.UserStatusActive,
		CreatedAt:   time.Now(),
	}
	return result
}

func (u UserBSON) ConvertToJSON() UserInfo {
	result := UserInfo{
		ID:          u.ID,
		FullName:    u.FullName,
		Email:       u.Email,
		Username:    u.Username,
		Avatar:      u.Avatar,
		Gender:      u.Gender,
		DateOfBirth: u.DateOfBirth,
		Phone:       u.Phone,
		Address:     u.Address,
	}
	return result
}

func (u UserUpdate) ConvertToBSON() UserInfoBSON {
	result := UserInfoBSON{
		FullName:    u.FullName,
		Email:       u.Email,
		Phone:       u.Phone,
		DateOfBirth: u.DateOfBirth,
		Gender:      u.Gender,
		Address:     u.Address,
	}
	return result
}

// Validate form body
func (body UserRegister) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.Username,
			validation.Required.Error("Username is required"),
			validation.Length(5, 30).Error("UserName is length: 5 -> 30"),
		),
		validation.Field(
			&body.Password,
			validation.Required.Error("Password is required"),
			validation.Length(5, 30).Error("Password is length: 5 -> 30"),
		),
		validation.Field(
			&body.Email,
			validation.Required.Error("Email is required"),
			is.Email.Error("Email invalidate"),
		),
		validation.Field(
			&body.FullName,
			validation.Required.Error("Fullname is required"),
			validation.Length(5, 30).Error("Fullname is length: 5 -> 30"),
		),
		validation.Field(
			&body.Gender,
			validation.Required.Error("Gender is required"),
			validation.In("female", "male"),
		),
		validation.Field(
			&body.DateOfBirth,
			validation.Required.Error("DateOfBirth is required"),
		),
	)
}

func (body UserLogin) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.Username,
			validation.Required.Error("Username is required"),
			validation.Length(5, 30).Error("UserName is length: 5 -> 30"),
		),
		validation.Field(
			&body.Password,
			validation.Required.Error("Password is required"),
			validation.Length(5, 30).Error("Password is length: 5 -> 30"),
		),
	)
}

func (body UserChangePassword) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.CurrentPassword,
			validation.Required.Error("CurrentPassword is required"),
			validation.Length(5, 30).Error("CurrentPassword is length: 5 -> 30"),
		),
		validation.Field(
			&body.NewPassword,
			validation.Required.Error("NewPassword is required"),
			validation.Length(5, 30).Error("NewPassword is length: 5 -> 30"),
		),
	)
}

func (body UserUpdate) Validate() error {
	return validation.ValidateStruct(&body,
		validation.Field(
			&body.FullName,
			validation.Required.Error("FullName is required"),
			validation.Length(5, 30).Error("FullName is length: 5 -> 30"),
		),
		validation.Field(
			&body.Email,
			validation.Required.Error("Email is required"),
			validation.Length(5, 30).Error("Email is length: 5 -> 30"),
		),
		validation.Field(
			&body.Phone,
			validation.Required.Error("Phone is required"),
			validation.Length(5, 30).Error("Phone is length: 5 -> 30"),
		),
		validation.Field(
			&body.DateOfBirth,
			validation.Required.Error("DateOfBirth is required"),
		),
		validation.Field(
			&body.Address,
			validation.Required.Error("Address is required"),
			validation.Length(5, 30).Error("Address is length: 5 -> 30"),
		),
	)
}
