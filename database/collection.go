package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	admins     = "admins"
	users      = "users"
	categories = "categories"
	brands     = "brands"
	products   = "products"
	reviews    = "reviews"
	replies    = "replies"
	carts      = "carts"
	cartItems  = "cartItems"
	orders     = "orders"
	orderItems = "orderItems"
	// ...
)

func UserCol() *mongo.Collection {
	return db.Collection(users)
}

func AdminCol() *mongo.Collection {
	return db.Collection(admins)
}

func ProductCol() *mongo.Collection {
	return db.Collection(products)
}

func CategoryCol() *mongo.Collection {
	return db.Collection(categories)
}

func BrandCol() *mongo.Collection {
	return db.Collection(brands)
}

func ReviewCol() *mongo.Collection {
	return db.Collection(reviews)
}

func ReplyCol() *mongo.Collection {
	return db.Collection(replies)
}

func OrderCol() *mongo.Collection {
	return db.Collection(orders)
}

func OrderItemCol() *mongo.Collection {
	return db.Collection(orderItems)
}
func CartCol() *mongo.Collection {
	return db.Collection(carts)
}
func CartItemCol() *mongo.Collection {
	return db.Collection(cartItems)
}

// ...
