package publisher

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListDeletedPublishers(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listDeletedPublishers")

	publishers := []models.PublisherResponse{}

	result, err := db.Query("SELECT * FROM PUBLISHERS WHERE DELETE_DT IS NOT NULL")

	if err != nil {
		logger.Errorf("error pulling deleted publishers: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	defer result.Close()

	for result.Next() {
		var publisher models.PublisherResponse
		if err := result.Scan(
			&publisher.ID,
			&publisher.Name,
			&publisher.CreateDate,
			&publisher.UpdateDate,
			&publisher.DeleteDate); err != nil {
			logger.Errorf("error populating publisher: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating publisher")
			return
		}

		publishers = append(publishers, publisher)
	}

	handler.SendSuccess(context, "list deleted publishers", publishers)

}
