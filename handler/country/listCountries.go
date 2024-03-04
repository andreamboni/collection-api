package country

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListCountriesHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listAuthors")

	countries := []models.CountryResponse{}

	result, err := db.Query("SELECT * FROM COUNTRIES")

	if err != nil {
		logger.Errorf("error pulling countries data: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error pulling countries data")
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
			logger.Errorf("error populating countries: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating country")
			return
		}

		countries = append(countries, country)
	}

	if len(countries) == 0 {
		handler.SendSuccess(context, "list countries", "there no countries in the database")
		return
	}

	handler.SendSuccess(context, "list countries", countries)
}
