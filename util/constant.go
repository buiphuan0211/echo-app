package util

import "context"

var (
	CreateSuccessFully            = "Create successfully"
	YouCannotEditThis             = "You cannot edit this"
	UpdateSuccessFully            = "Update successfully"
	UserNotFound                  = "User not found"
	DeleteSuccessFully            = "Delete successfully"
	InvalidData                   = "Invalid data"
	EmailOrPasswordWrong          = "Email or password wrong"
	EmailOrUsernameIsAlreadyExist = "Email or username already exists"
	LoginSuccessFully             = "Login successfully"
	InvalidToken                  = "Invalid token"

	Ctx = context.Background()

	// Auth
	NotExistUser        = "user not existed in database"
	WrongPassword       = "Wrong password"
	GenerateTokenFailed = "Generate token failed"

	// Status User
	UserStatusActive  = "ACTIVE"
	UserStatusBlocked = "BLOCKED"

	//
	CurrentPasswordIsIncorrect = "CurrentPassword is incorrect"

	// Order
	OrderStatusPending    = "PENDING"
	OrderStatusConfirmed  = "CONFIRMED"
	OrderStatusDelivering = "DELIVERING"
	OrderStatusDelivered  = "DELIVERED"
	OrderStatusCancel     = "CANCEL"

	// Category
	CategoryStatusEnabled  = "ENABLE"
	CategoryStatusDisabled = "DISABLE"

	// Product
	ProductStatusEnabled  = "ENABLE"
	ProductStatusDisabled = "DISABLE"
)
