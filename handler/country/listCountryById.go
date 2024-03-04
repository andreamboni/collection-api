package country

import (
	"database/sql"
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListCountryByIdHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listCountryById")

	country := models.CountryResponse{}

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	row := db.QueryRow("SELECT * FROM COUNTRIES WHERE ID = ?", id)

	if err := row.Scan(
		&country.ID,
		&country.Name,
		&country.CreateDate,
		&country.UpdateDate,
		&country.DeleteDate); err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("no country was found: %v", err.Error())
			handler.SendError(context, http.StatusNotFound, "no country was found")
			return
		}

		logger.Errorf("error populating country: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error populating country")
	}

	handler.SendSuccess(context, "list country by id", country)
}
