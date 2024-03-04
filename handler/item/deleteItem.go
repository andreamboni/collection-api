package item

import (
	"net/http"
	"time"

	"collection.com/config"
	"collection.com/handler"
	"github.com/gin-gonic/gin"
)

func DeleteItemHandler(context *gin.Context) {

	db := config.GetMySQL()
	logger := config.GetLogger("deleteItem")

	id := context.Query("id")

	if id == "" {
		logger.Error("no id was informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	result, err := db.Exec("UPDATE ITEMS SET DELETE_DT = ? WHERE ID = ?", time.Now(), id)

	if err != nil {
		logger.Errorf("error deleting item: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("error deleting item: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		logger.Error("error deleting item. no item was found")
		handler.SendError(context, http.StatusNotFound, "no item was found with id "+id)
		return
	}

	handler.SendSuccess(context, "item deleted", result)

}
