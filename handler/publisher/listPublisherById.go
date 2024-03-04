package publisher

import (
	"database/sql"
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListPublisherByIdHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listPublisherById")

	publisher := models.PublisherResponse{}

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	row := db.QueryRow("SELECT * FROM PUBLISHERS WHERE ID = ?", id)

	if err := row.Scan(
		&publisher.ID,
		&publisher.Name,
		&publisher.CreateDate,
		&publisher.UpdateDate,
		&publisher.DeleteDate); err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("no publisher was found: %v", err.Error())
			handler.SendError(context, http.StatusNotFound, "no publisher was found")
			return
		}

		logger.Errorf("error populating publisher: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error populating publisher")
	}

	handler.SendSuccess(context, "list publisher by id", publisher)
}
