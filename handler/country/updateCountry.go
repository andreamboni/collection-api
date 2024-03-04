package country

import (
	"net/http"
	"strconv"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func UpdateCountryHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("updateCountry")

	request := CreateCountryRequest{}
	context.BindJSON(&request)

	id := context.Query("id")

	if id == "" {
		logger.Error("id was not informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	country := models.Country{
		Name: request.Name,
	}

	_, err := db.Exec("UPDATE COUNTRIES SET NAME = ? WHERE ID = ?", country.Name, id)

	if err != nil {
		logger.Errorf("error updating country: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	countryId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		logger.Error("id is not a number")
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	country.ID = countryId

	handler.SendSuccess(context, "country updated", country)
}
