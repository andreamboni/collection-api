package country

import (
	"net/http"
	"time"

	"collection.com/config"
	"collection.com/handler"
	"github.com/gin-gonic/gin"
)

func DeleteCountryHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("deleteCountry")

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	result, err := db.Exec("UPDATE COUNTRIES SET DELETE_DT = ? WHERE ID = ?", time.Now(), id)

	if err != nil {
		logger.Errorf("error deleting country: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("error deleting country: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		logger.Error("error deleting country. no country was found")
		handler.SendError(context, http.StatusNotFound, "no country was found with id "+id)
		return
	}

	handler.SendSuccess(context, "country deleted", "country deleted")
}
