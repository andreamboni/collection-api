package author

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"github.com/gin-gonic/gin"
)

func RecoverAuthorHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("recoverAuthor")

	id := context.Query("id")

	if id == "" {
		logger.Error("id was not informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	result, err := db.Exec("UPDATE AUTHORS SET DELETE_DT = ? WHERE ID = ?", nil, id)

	if err != nil {
		logger.Errorf("error recovering author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("error recovering author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		logger.Error("error recovering author. no author was found")
		handler.SendError(context, http.StatusNotFound, "no author was found with id "+id)
		return
	}

	handler.SendSuccess(context, "recovered author", "recovered author")
}
