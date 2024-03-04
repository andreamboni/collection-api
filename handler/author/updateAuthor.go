package author

import (
	"net/http"
	"strconv"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func UpdateAuthorHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("updateAuthor")

	request := CreateAuthorRequest{}
	context.BindJSON(&request)

	id := context.Query("id")

	if id == "" {
		logger.Error("id was not informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	author := models.Author{
		Name:        request.Name,
		Nationality: request.Nationality,
	}

	_, err := db.Exec("UPDATE AUTHORS SET NAME = ?, NATIONALITY = ? WHERE ID = ?", author.Name, author.Nationality, id)

	if err != nil {
		logger.Errorf("error updating author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	authorId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		logger.Error("id is not a number")
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	author.ID = authorId

	handler.SendSuccess(context, "author updated", author)
}
