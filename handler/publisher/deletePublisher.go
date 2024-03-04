package publisher

import (
	"net/http"
	"time"

	"collection.com/config"
	"collection.com/handler"
	"github.com/gin-gonic/gin"
)

func DeletePublisherHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("deletePublisher")

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	result, err := db.Exec("UPDATE PUBLISHERS SET DELETE_DT = ? WHERE ID = ?", time.Now(), id)

	if err != nil {
		logger.Errorf("error deleting publisher: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("error deleting publisher: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		logger.Error("error deleting publisher. no publisher was found")
		handler.SendError(context, http.StatusNotFound, "no publisher was found with id "+id)
		return
	}

	handler.SendSuccess(context, "publisher deleted", "publisher deleted")
}
