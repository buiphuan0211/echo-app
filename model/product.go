package model

import (
	"echo-app/util"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ProductBSON struct {
		Id          primitive.ObjectID `bson:"_id"`
		Name        string             `bson:"name"`
		CategoryId  primitive.ObjectID `bson:"categoryId"`
		BrandId     primitive.ObjectID `bson:"brandId"`
		Price       uint               `bson:"price"`
		Description string             `bson:"description"`
		Images      []string           `bson:"images"`
		Quantity    uint               `bson:"quantity"`
		Rest        uint               `bson:"rest"`
		Status      string             `bson:"status"`
		CreatedAt   time.Time          `bson:"createdAt"`
		UpdatedAt   time.Time          `bson:"updatedAt"`
	}

	ProductQuery struct {
		Page       int64  `query:"page"`
		Name       string `query:"name"`
		CategoryId string `query:"categoryId"`
		PriceFrom  string `query:"priceFrom"`
		BrandId    string `query:"brandId"`
		Sort       string `query:"sort"`
	}

	ProductResponse struct {
		Id          primitive.ObjectID     `json:"_id"`
		Name        string                 `json:"name"`
		Category    map[string]interface{} `json:"category"`
		Brand       map[string]interface{} `json:"brand"`
		Price       uint                   `json:"price"`
		Description string                 `json:"description"`
		Images      []string               `json:"images"`
		Quantity    uint                   `json:"quantity"`
		Rest        uint                   `json:"rest"`
		Status      string                 `json:"status"`
		CreatedAt   string                 `json:"createdAt"`
		UpdatedAt   string                 `json:"updatedAt"`
	}

	ProductCreate struct {
		Name        string             `json:"name"`
		CategoryId  primitive.ObjectID `json:"categoryId"`
		BrandId     primitive.ObjectID `json:"brandId"`
		Price       uint               `json:"price"`
		Description string             `json:"description"`
		Images      []string           `json:"images"`
		Quantity    uint               `json:"quantity"`
		Rest        uint               `json:"rest"`
		Status      string             `json:"status"`
	}

	ProductUpdate struct {
		Name        string             `json:"name"`
		CategoryId  primitive.ObjectID `json:"categoryId"`
		BrandId     primitive.ObjectID `json:"brandId"`
		Price       uint               `json:"price"`
		Quantity    uint               `json:"quantity"`
		Rest        uint               `json:"rest"`
		Images      []string           `json:"images"`
		Description string             `json:"description"`
	}
	ProductUpdateBSON struct {
		Name        string             `bson:"name"`
		CategoryId  primitive.ObjectID `bson:"categoryId"`
		BrandId     primitive.ObjectID `bson:"brandId"`
		Price       uint               `bson:"price"`
		Quantity    uint               `bson:"quantity"`
		Rest        uint               `bson:"rest"`
		Images      []string           `bson:"images"`
		Description string             `bson:"description"`
	}
)

func (p ProductCreate) ConvertToBSON() ProductBSON {
	result := ProductBSON{
		Id:          primitive.NewObjectID(),
		Name:        p.Name,
		CategoryId:  p.CategoryId,
		BrandId:     p.BrandId,
		Price:       p.Price,
		Description: p.Description,
		Images:      p.Images,
		Quantity:    p.Quantity,
		Rest:        p.Rest,
		Status:      util.ProductStatusEnabled,
	}
	return result
}

func (p ProductUpdate) ConvertToBSON() ProductUpdateBSON {
	result := ProductUpdateBSON{
		Name:        p.Name,
		CategoryId:  p.CategoryId,
		BrandId:     p.BrandId,
		Price:       p.Price,
		Description: p.Description,
		Images:      p.Images,
		Quantity:    p.Quantity,
		Rest:        p.Rest,
	}
	return result
}

func (p ProductBSON) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.Images, validation.Length(1, 20)),
	)
}
