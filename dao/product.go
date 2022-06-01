package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
	"fmt"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var limit int64 = 10

type Product struct{}

func (p Product) GetListProduct(query model.ProductQuery) ([]model.ProductResponse, error) {
	var (
		listProduct []model.ProductResponse
		sort        = bson.M{}
		filter      = bson.M{"status": "enable"}
	)

	//init filter
	if query.CategoryId != "" {
		filter["categoryId"], _ = primitive.ObjectIDFromHex(query.CategoryId)
	}
	if query.BrandId != "" {
		filter["brandId"], _ = primitive.ObjectIDFromHex(query.BrandId)
	}
	if query.Name != "" {
		filter["name"] = query.Name
	}
	if query.PriceFrom != "" {
		priceFrom, _ := strconv.ParseInt(query.PriceFrom, 10, 64)
		filter["price"] = bson.M{
			"$lte": priceFrom,
		}
	}

	//init sort
	if query.Sort != "" {
		sortArr := strings.Split(query.Sort, ",")
		for _, v := range sortArr {
			var value int
			if string([]rune(v)[1]) != "-" {
				fmt.Println(string([]rune(v)[0]))
				value = -1
			} else {
				value = 1
			}
			sort[v] = value
		}
	}

	//init pipeline
	pipeline := []bson.M{
		{"$match": filter},
		{"$lookup": bson.M{
			"from":         "brands",
			"localField":   "brandId",
			"foreignField": "_id",
			"as":           "brand",
		}},
		{"$unwind": "$brand"},
		{"$lookup": bson.M{
			"from":         "categories",
			"localField":   "categoryId",
			"foreignField": "_id",
			"as":           "category",
		}},
		{"$unwind": "$category"},
		{"$skip": query.Page * limit},
		{"$limit": limit},
		{"sort": sort},
	}

	cursor, err := database.ProductCol().Aggregate(util.Ctx, pipeline)
	if err != nil {
		return listProduct, err
	}

	err = cursor.All(util.Ctx, &listProduct)

	fmt.Println("LIST PRODUCT IN DAO: ", listProduct)

	return listProduct, nil
}

func (p Product) GetProductDetail(id primitive.ObjectID) (*model.ProductResponse, error) {
	var (
		results *model.ProductResponse
		filter  = bson.M{"_id": id}
	)

	err := database.ProductCol().FindOne(util.Ctx, filter).Decode(&results)
	return results, err
}

func (p Product) Create(body model.ProductBSON) (string, error) {
	result, err := database.ProductCol().InsertOne(util.Ctx, body)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (p Product) Update(id primitive.ObjectID, body model.ProductUpdateBSON) error {
	_, err := database.ProductCol().UpdateOne(util.Ctx, bson.M{"_id": id}, bson.M{"$set": body})
	if err != nil {
		return err
	}

	return nil
}

func (p Product) UpdateStatus(id primitive.ObjectID, status string) error {
	_, err := database.ProductCol().UpdateOne(util.Ctx, bson.M{"_id": id}, bson.D{
		{"$set", bson.D{{"status", status}}},
	})
	if err != nil {
		return err
	}

	return nil
}

func (p Product) GetByID(id primitive.ObjectID) (model.ProductBSON, error) {
	var (
		product model.ProductBSON
	)

	filter := bson.M{"_id": id}

	if err := database.ProductCol().FindOne(util.Ctx, filter).Decode(&product); err != nil {
		return product, err
	}

	return product, nil
}
