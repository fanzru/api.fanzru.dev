package controllers

import (
	"api.fanzru.dev/models"
	"go.mongodb.org/mongo-driver/bson"

	//"api.fanzru.dev/utils"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllMessage(c echo.Context)  error{
	return c.String(http.StatusOK, "Sayang sama dia")
}

func Message(c echo.Context) error{
	result := []models.Post{}
	err := mgm.Coll(&models.Post{}).SimpleFind(&result, bson.M{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "cie error"} )
	}
	return c.JSON(http.StatusOK, result)
}

func SearchMessage(c echo.Context) error{
	type reqData struct {
		Name    	string `json:"name" bson:"name"`
	}
	var data reqData
	if err := c.Bind(&data); err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "cie error"} )
	}
	result := []models.Post{}
	err := mgm.Coll(&models.Post{}).SimpleFind(&result, bson.M{"name": data.Name})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "cie error"} )
	}
	return c.JSON(http.StatusOK, result)

}

func CreateMessage(c echo.Context) error {
	type reqData struct {
		Name    	string `json:"name" bson:"name"`
		Message		string `json:"message" bson:"message"`
	}
	var data reqData

	if err := c.Bind(&data); err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "cie error"} )
	}

	post := &models.Post{
		Name:      data.Name,
		Message:    data.Message,
	}
	//utils.PrintMaptoJSON(post)
	// To get model's collection, just call to mgm.Coll() method.
	err := mgm.Coll(post).Create(post)
	if err != nil {
		// Log your error
		return c.JSON(http.StatusInternalServerError,internalError)
	}
	return c.JSON(http.StatusOK, post)
}
