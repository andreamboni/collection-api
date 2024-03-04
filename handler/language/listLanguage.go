package language

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListLanguageHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listLanguage")

	languages := []models.LanguageResponse{}

	result, err := db.Query("SELECT * FROM LANGUAGES")

	if err != nil {
		logger.Errorf("error pulling the language data: %v", err.Error())
		handler.SendError(context, http.StatusNotFound, "error pulling the language data")
		return
	}

	defer result.Close()

	for result.Next() {
		var language models.LanguageResponse
		if err := result.Scan(&language.ID,
			&language.Name,
			&language.CreateDate,
			&language.UpdateDate,
			&language.DeleteDate); err != nil {
			logger.Errorf("error populating language: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating language")
			return
		}
		languages = append(languages, language)
	}

	handler.SendSuccess(context, "list languages", languages)
}
