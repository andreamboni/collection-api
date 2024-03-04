package language

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func CreateLanguageHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("createLanguage")

	request := CreateLanguageRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	language := models.Language{
		Name: request.Name,
	}

	result, err := db.Exec("INSERT INTO LANGUAGES(NAME) VALUES(?)", language.Name)

	if err != nil {
		logger.Errorf("error creating language: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Errorf("error creating language: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	language.ID = id

	handler.SendSuccess(context, "language created", language)
}
