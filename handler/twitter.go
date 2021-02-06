package handler

import (
	"net/http"
	"time"

	// "encoding/json"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	database "github.com/klan300/exceed17/database"
	
)

type TwitInput  struct {
	PublishAt time.Time `json:"publish_At"`
	Author string `json:"author"`
	Content string `json:"content"`
}

func GetTweet(request echo.Context) error {

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("tweet")
	defer db.Client().Disconnect(ctx)

	options := options.Find()
	options.SetSort(bson.D{{"publishat", -1}})
	options.SetLimit(10)

	var twitInput []*TwitInput

	cursor, err := collection.Find(ctx, bson.D{},options)
	
	if err != nil {
		return request.JSON(http.StatusInternalServerError, err)
	}

	cursor.All(ctx,&twitInput)

	return request.JSON(http.StatusOK, twitInput)
}


func PostTweet(request echo.Context) error {

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("tweet")
	defer db.Client().Disconnect(ctx)

	var twitInput TwitInput
	err :=  request.Bind(&twitInput)

	twitInput.PublishAt = time.Now()

	if err != nil {
		return request.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	result ,err := collection.InsertOne(ctx, twitInput)

	if err != nil {
		return request.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return request.JSON(http.StatusOK, result)
}


