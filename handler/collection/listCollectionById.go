package collection

import (
	"database/sql"
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListCollectionByIdHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listCollectionById")

	collection := models.CollectionResponse{}

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	result := db.QueryRow("SELECT * FROM COLLECTIONS WHERE ID = ?", id)

	if err := result.Scan(
		&collection.ID,
		&collection.Name,
		&collection.CreateDate,
		&collection.UpdateDate,
		&collection.DeleteDate); err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("no collection was found: %v", err.Error())
			handler.SendError(context, http.StatusNotFound, "no collection was found")
			return
		}

		logger.Errorf("error populating collection: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error populating collection")
	}

	handler.SendSuccess(context, "list collection by id", collection)
}
