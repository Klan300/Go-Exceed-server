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

func GetDataById(request echo.Context) error {

	id := request.Param("studentId")

	var students Student

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("exceedFrontend")
	defer db.Client().Disconnect(ctx)

	err := collection.FindOne(ctx, bson.M{"id":id}).Decode(&students)

	if err != nil {
		return request.JSON(http.StatusNotFound, echo.NotFoundHandler)
	}

	return request.JSON(http.StatusOK, students)
}


func PutDataById(request echo.Context) error {

	id := request.Param("studentId")

	var students Student

	err :=  request.Bind(&students)

	if err != nil {
		return request.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	if students.Id != id {
		return request.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	} 

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("exceedFrontend")
	defer db.Client().Disconnect(ctx)

	_ ,err  = collection.ReplaceOne(ctx, bson.M{"id":id}, students)

	if err != nil {
		return request.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return request.JSON(http.StatusOK, students)
}

func PatchDataById(request echo.Context) error {

	id := request.Param("studentId")

	var students Student

	err :=  request.Bind(&students)

	if err != nil {
		return request.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	if students.Id != "" || students.Question != "" {
		return request.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	ctx, db := database.DatabaseConnect()
	collection := db.Collection("exceedFrontend")
	defer db.Client().Disconnect(ctx)

	_ , err = collection.UpdateOne(ctx, bson.M{"id":id}, bson.D{{"$set", bson.D{{"answer", students.Answer}}}})

	if err != nil {
		return request.JSON(http.StatusNotFound, echo.NotFoundHandler)
	}

	err	= collection.FindOne(ctx, bson.M{"id":id}).Decode(&students)

	if err != nil {
		return request.JSON(http.StatusNotFound, echo.NotFoundHandler)
	}

	return request.JSON(http.StatusOK, students)
}
