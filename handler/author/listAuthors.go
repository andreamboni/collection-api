package author

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListAuthorsHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listAuthors")

	authors := []models.AuthorResponse{}

	result, err := db.Query("SELECT * FROM AUTHORS")

	if err != nil {
		logger.Errorf("error pulling authors data: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error pulling authors data")
		return
	}

	defer result.Close()

	for result.Next() {
		var author models.AuthorResponse
		if err := result.Scan(
			&author.ID,
			&author.Name,
			&author.Nationality,
			// &author.BirthDate,
			// &author.DeathDate,
			&author.CreateDate,
			&author.UpdateDate,
			&author.DeleteDate); err != nil {
			logger.Errorf("error populating authors: %v", err.Error())
			handler.SendError(context, http.StatusInternalServerError, "error populating author")
			return
		}

		authors = append(authors, author)
	}

	if len(authors) == 0 {
		handler.SendSuccess(context, "list authors", "there no authors in the database")
		return
	}

	handler.SendSuccess(context, "list authors", authors)
}
