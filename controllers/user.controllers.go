package controllers
import (
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"api.fanzru.dev/models"
	"net/http"
)

type M map[string]interface{}

// Define our errors:
var internalError = M{"message": "internal error"}

func GetUser(c echo.Context)  error{
	// User ID from path `users/:id`
	return c.String(http.StatusOK, "Sayang sama dia")
}

// Create handler create new book.
func Create(c echo.Context) error {
	type reqData struct {
		Name      string `json:"name" bson:"name"`
		Author    string `json:"author" bson:"author"`
	}
	var data reqData

	if err := c.Bind(&data); err != nil{
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "cie error"} )
	}

	user := &models.User{
		Name:      data.Name,
		Author:    data.Author,
	}

	// To get model's collection, just call to mgm.Coll() method.
	err := mgm.Coll(user).Create(user)
	if err != nil {
		// Log your error
		return c.JSON(http.StatusInternalServerError,internalError)
	}
	return c.JSON(http.StatusOK, user)
}
