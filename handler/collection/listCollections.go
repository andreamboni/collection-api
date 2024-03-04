package collection

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListCollectionsHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listAuthors")

	collections := []models.CollectionResponse{}

	result, err := db.Query("SELECT * FROM COLLECTIONS")

	if err != nil {
		logger.Errorf("error pulling collections data: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error pulling collections data")
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
			logger.Errorf("error populating collections: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating collection")
			return
		}

		collections = append(collections, collection)
	}

	if len(collections) == 0 {
		handler.SendSuccess(context, "list collections", "there no collections in the database")
		return
	}

	handler.SendSuccess(context, "list collections", collections)
}
