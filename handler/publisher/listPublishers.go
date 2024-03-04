package publisher

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListPublishersHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listPublishers")

	publishers := []models.PublisherResponse{}

	result, err := db.Query("SELECT * FROM PUBLISHERS")

	if err != nil {
		logger.Errorf("error pulling publishers data: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error pulling publishers data")
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
			logger.Errorf("error populating publishers: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating publisher")
			return
		}

		publishers = append(publishers, publisher)
	}

	if len(publishers) == 0 {
		handler.SendSuccess(context, "list publishers", "there no publishers in the database")
		return
	}

	handler.SendSuccess(context, "list publishers", publishers)
}
