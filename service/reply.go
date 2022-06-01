package service

import (
	"echo-app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reply struct{}

func (Reply) CreateReply(userId string, reviewId string, body model.CreateReply) (string, error) {
	userOjbID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", err
	}

	reviewOjbID, err := primitive.ObjectIDFromHex(reviewId)
	if err != nil {
		return "", err
	}

	//init insert data
	insertData := model.Reply{
		UserId:    userOjbID,
		Content:   body.Content,
		ReviewId:  reviewOjbID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insertedId, err := replyDAO.CreateReply(insertData)
	return insertedId, err
}

func (Reply) UpdateReply(userId string, replyId string, body model.CreateReply) (*mongo.UpdateResult, error) {
	userOjbID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	replyOjbID, err := primitive.ObjectIDFromHex(replyId)
	if err != nil {
		return nil, err
	}

	updateData := model.Reply{
		Content:   body.Content,
		UpdatedAt: time.Now(),
	}

	results, err := replyDAO.UpdateReply(userOjbID, replyOjbID, updateData)
	return results, err
}

func (Reply) DeleteReply(userId string, replyId string) (*mongo.DeleteResult, error) {
	userOjbID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	replyOjbID, err := primitive.ObjectIDFromHex(replyId)
	if err != nil {
		return nil, err
	}

	results, err := replyDAO.DeleteReply(userOjbID, replyOjbID)
	return results, err
}
