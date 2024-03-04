package author

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func CreateAuthorHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("createAuthor")

	request := CreateAuthorRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	author := models.Author{
		Name:        request.Name,
		Nationality: request.Nationality,
		// BirthDate:   birthDate,
		// DeathDate:   deathDate,
	}

	result, err := db.Exec("INSERT INTO AUTHORS (NAME, NATIONALITY) VALUES (?, ?)", author.Name, author.Nationality)

	if err != nil {
		logger.Errorf("error creating author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Errorf("error creating author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	author.ID = id

	handler.SendSuccess(context, "author created", author)
}
