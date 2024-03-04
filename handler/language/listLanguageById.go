package language

import (
	"database/sql"
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListLanguageByIdHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listLanguageById")

	language := models.LanguageResponse{}

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	row := db.QueryRow("SELECT * FROM LANGUAGES WHERE ID = ?", id)

	if err := row.Scan(
		&language.ID,
		&language.Name,
		&language.CreateDate,
		&language.UpdateDate,
		&language.DeleteDate); err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("no language was found: %v", err.Error())
			handler.SendError(context, http.StatusNotFound, "no language was found")
			return
		}

		logger.Errorf("error populating language: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error populating language")
	}

	handler.SendSuccess(context, "list language by id", language)
}
