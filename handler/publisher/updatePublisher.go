package publisher

import (
	"net/http"
	"strconv"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func UpdatePublisherHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("updatePublisher")

	request := CreatePublisherRequest{}
	context.BindJSON(&request)

	id := context.Query("id")

	if id == "" {
		logger.Error("id was not informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	publisher := models.Publisher{
		Name: request.Name,
	}

	_, err := db.Exec("UPDATE PUBLISHERS SET NAME = ? WHERE ID = ?", publisher.Name, id)

	if err != nil {
		logger.Errorf("error updating publisher: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	publisherId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		logger.Error("id is not a number")
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	publisher.ID = publisherId

	handler.SendSuccess(context, "publisher updated", publisher)
}
