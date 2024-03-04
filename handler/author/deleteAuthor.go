package author

import (
	"net/http"
	"time"

	"collection.com/config"
	"collection.com/handler"
	"github.com/gin-gonic/gin"
)

func DeleteAuthorHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("deleteAuthor")

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	result, err := db.Exec("UPDATE AUTHORS SET DELETE_DT = ? WHERE ID = ?", time.Now(), id)

	if err != nil {
		logger.Errorf("error deleting author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("error deleting author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		logger.Error("error deleting author. no author was found")
		handler.SendError(context, http.StatusNotFound, "no author was found with id "+id)
		return
	}

	handler.SendSuccess(context, "author deleted", "author deleted")
}
