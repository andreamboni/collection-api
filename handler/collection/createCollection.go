package collection

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func CreateCollectionHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("createCollection")

	request := CreateCollectionRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	collection := models.Collection{
		Name: request.Name,
	}

	result, err := db.Exec("INSERT INTO COLLECTIONS (NAME) VALUES (?)", collection.Name)

	if err != nil {
		logger.Errorf("error creating collection: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Errorf("error creating collection: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	collection.ID = id

	handler.SendSuccess(context, "collection created", collection)
}
