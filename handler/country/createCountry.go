package country

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func CreateCountryHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("createCountries")

	request := CreateCountryRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	country := models.Country{
		Name: request.Name,
	}

	result, err := db.Exec("INSERT INTO COUNTRIES (NAME) VALUES (?)", country.Name)

	if err != nil {
		logger.Errorf("error creating country: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Errorf("error creating country: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	country.ID = id

	handler.SendSuccess(context, "country created", country)
}
