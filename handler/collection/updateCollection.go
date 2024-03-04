package collection

import (
	"net/http"
	"strconv"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func UpdateCollectionHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("updateCollection")

	request := CreateCollectionRequest{}
	context.BindJSON(&request)

	id := context.Query("id")

	if id == "" {
		logger.Error("id was not informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	collection := models.Collection{
		Name: request.Name,
	}

	_, err := db.Exec("UPDATE COLLECTIONS SET NAME = ? WHERE ID = ?", collection.Name, id)

	if err != nil {
		logger.Errorf("error updating collection: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	collectionId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		logger.Error("id is not a number")
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	collection.ID = collectionId

	handler.SendSuccess(context, "collection updated", collection)
}
