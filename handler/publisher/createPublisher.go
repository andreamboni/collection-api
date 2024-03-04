package publisher

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func CreatePublisherHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("createPublisher")

	request := CreatePublisherRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	publisher := models.Publisher{
		Name: request.Name,
	}

	result, err := db.Exec("INSERT INTO PUBLISHERS (NAME) VALUES (?)", publisher.Name)

	if err != nil {
		logger.Errorf("error creating publisher: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Errorf("error creating publisher: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	publisher.ID = id

	handler.SendSuccess(context, "publisher created", publisher)
}
