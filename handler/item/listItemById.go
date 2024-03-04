package item

import (
	"database/sql"
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListItemByIdHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listItemById")

	item := models.ItemResponse{}

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	row := db.QueryRow("SELECT * FROM ITEMS WHERE ID = ?", id)

	if err := row.Scan(
		&item.ID,
		&item.Collection,
		&item.Title,
		&item.Author,
		&item.Publisher,
		&item.ItemType,
		&item.ItemFormat,
		&item.PagesNumber,
		&item.Edition,
		&item.EditionYear,
		&item.Binding,
		&item.Language,
		&item.Country,
		&item.Copies,
		&item.CreateDate,
		&item.UpdateDate,
		&item.DeleteDate); err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("no item was found: %v", err.Error())
			handler.SendError(context, http.StatusNotFound, "no item was found")
			return
		}

		logger.Errorf("error populating item: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error populating item")
		return
	}

	handler.SendSuccess(context, "list-item-by-id", item)

}
