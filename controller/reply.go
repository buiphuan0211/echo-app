package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type Reply struct{}

func (Reply) CreateReply(c echo.Context) error {
	reviewId := c.Param("id")

	payloadInterface := c.Get("payload")
	payload := payloadInterface.(model.CreateReply)

	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	InsertedId, err := replyService.CreateReply(userId, reviewId, payload)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, InsertedId, util.UpdateSuccessFully)
}

func (Reply) UpdateReply(c echo.Context) error {
	replyId := c.Param("id")

	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	body := c.Get("body").(model.CreateReply)

	results, err := replyService.UpdateReply(userId, replyId, body)
	if err != nil || results.MatchedCount == 0 {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, replyId, util.UpdateSuccessFully)
}

func (Reply) DeleteReply(c echo.Context) error {
	replyId := c.Param("id")

	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	results, err := replyService.DeleteReply(userId, replyId)
	if err != nil || results.DeletedCount == 0 {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, nil, util.DeleteSuccessFully)
}
