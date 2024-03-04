package collection

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListDeletedCollection(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listDeletedCollection")

	collections := []models.CollectionResponse{}

	result, err := db.Query("SELECT * FROM COLLECTIONS WHERE DELETE_DT IS NOT NULL")

	if err != nil {
		logger.Errorf("error pulling deleted collections: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	defer result.Close()

	for result.Next() {
		var collection models.CollectionResponse
		if err := result.Scan(
			&collection.ID,
			&collection.Name,
			&collection.CreateDate,
			&collection.UpdateDate,
			&collection.DeleteDate); err != nil {
			logger.Errorf("error populating collection: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating collection")
			return
		}

		collections = append(collections, collection)
	}

	handler.SendSuccess(context, "list deleted collections", collections)

}
