package country

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListDeletedCountries(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listDeletedCountries")

	countries := []models.CountryResponse{}

	result, err := db.Query("SELECT * FROM COUNTRIES WHERE DELETE_DT IS NOT NULL")

	if err != nil {
		logger.Errorf("error pulling deleted countries: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	defer result.Close()

	for result.Next() {
		var country models.CountryResponse
		if err := result.Scan(
			&country.ID,
			&country.Name,
			&country.CreateDate,
			&country.UpdateDate,
			&country.DeleteDate); err != nil {
			logger.Errorf("error populating country: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating country")
			return
		}

		countries = append(countries, country)
	}

	handler.SendSuccess(context, "list deleted countries", countries)

}
