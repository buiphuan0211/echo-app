package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Review struct{}

func (Review) GetListReview(ojbID primitive.ObjectID, query model.ReviewQuery) ([]model.ReviewResponse, error) {
	var listReview []model.ReviewResponse

	filter := bson.M{"productId": ojbID}

	if query.Rating != "" {
		rating, _ := strconv.ParseInt(query.Rating, 10, 64)
		filter["rating"] = rating
	}

	opts := new(options.FindOptions)
	opts.SetSkip(query.Page * limit)
	opts.SetLimit(limit)

	if query.Sort != "" {
		var value int
		if string([]rune(query.Sort)[0]) != "-" {
			value = -1
		} else {
			value = 1
		}
		sortMap := map[string]interface{}{
			"price": value,
		}
		opts.SetSort(sortMap)
	}
	cursor, err := database.ReviewCol().Find(util.Ctx, filter, opts)
	if err != nil {
		return listReview, err
	}

	err = cursor.All(util.Ctx, &listReview)
	return listReview, err
}

func (Review) CreateReview(insertData model.ReviewBSON) (*mongo.InsertOneResult, error) {
	result, err := database.ReviewCol().InsertOne(util.Ctx, insertData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
