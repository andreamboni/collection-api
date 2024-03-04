package author

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListDeletedAuthors(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listDeletedAuthors")

	authors := []models.AuthorResponse{}

	result, err := db.Query("SELECT * FROM AUTHORS WHERE DELETE_DT IS NOT NULL")

	if err != nil {
		logger.Errorf("error pulling deleted authors: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	defer result.Close()

	for result.Next() {
		var author models.AuthorResponse
		if err := result.Scan(
			&author.ID,
			&author.Name,
			&author.Nationality,
			&author.CreateDate,
			&author.UpdateDate,
			&author.DeleteDate); err != nil {
			logger.Errorf("error populating author: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating author")
			return
		}

		authors = append(authors, author)
	}

	handler.SendSuccess(context, "list deleted authors", authors)

}
