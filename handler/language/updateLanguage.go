package language

import (
	"net/http"
	"strconv"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func UpdateLanguageHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("updateLanguage")

	request := CreateLanguageRequest{}
	context.BindJSON(&request)

	id := context.Query("id")

	if id == "" {
		logger.Error("id was not informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	language := models.Language{
		Name: request.Name,
	}

	_, err := db.Exec("UPDATE LANGUAGES SET NAME = ? WHERE ID = ?", language.Name, id)

	if err != nil {
		logger.Errorf("error updating language: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	languageId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		logger.Error("id is not a number")
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	language.ID = languageId

	handler.SendSuccess(context, "language updated", language)
}
