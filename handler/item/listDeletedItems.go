package item

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListDeletedItems(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listDeletedItems")

	items := []models.ItemResponse{}

	result, err := db.Query("SELECT * FROM ITEMS WHERE DELETE_DT IS NOT NULL")

	if err != nil {
		logger.Errorf("error pulling deleted items: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	defer result.Close()

	for result.Next() {
		var item models.ItemResponse
		if err := result.Scan(
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
			logger.Errorf("error populating items: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating item")
			return
		}

		items = append(items, item)
	}

	handler.SendSuccess(context, "list deleted items", items)

}
