package handler

import (
	"net/http"

	// "encoding/json"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	
	database "github.com/klan300/exceed17/database"
	
)

type Student struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type answer struct {
	Answer   string `json:"answer"`
}


func GetDataById(request echo.Context) error {

	id := request.Param("studentId")

	var students Student

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("exceedFrontend")
	defer db.Client().Disconnect(ctx)

	err := collection.FindOne(ctx, bson.M{"id":id}).Decode(&students)

	if err != nil {
		return request.JSON(http.StatusInternalServerError, err)
	}

	return request.JSON(http.StatusOK, students)
}


func PutDataById(request echo.Context) error {

	id := request.Param("studentId")

	var students Student

	err :=  request.Bind(&students)

	if err != nil {
		return request.NoContent(http.StatusBadRequest)
	}

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("exceedFrontend")
	defer db.Client().Disconnect(ctx)

	status,err := collection.ReplaceOne(ctx, bson.M{"id":id}, students)

	if err != nil {
		return request.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return request.JSON(http.StatusOK, status)
}

func PatchDataById(request echo.Context) error {

	id := request.Param("studentId")

	var ans answer

	err :=  request.Bind(&ans)

	if err != nil {
		return request.NoContent(http.StatusBadRequest)
	}

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("exceedFrontend")
	defer db.Client().Disconnect(ctx)

	status,err := collection.UpdateOne(ctx, bson.M{"id":id}, bson.D{{"$set", bson.D{{"answer", ans.Answer}}}})

	if err != nil {
		return request.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return request.JSON(http.StatusOK, status)
}
