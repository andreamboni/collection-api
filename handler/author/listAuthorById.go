package author

import (
	"database/sql"
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func ListAuthorByIdHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("listAuthorById")

	author := models.AuthorResponse{}

	id := context.Query("id")

	if id == "" {
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
		return
	}

	row := db.QueryRow("SELECT * FROM AUTHORS WHERE ID = ?", id)

	if err := row.Scan(
		&author.ID,
		&author.Name,
		&author.Nationality,
		&author.CreateDate,
		&author.UpdateDate,
		&author.DeleteDate); err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("no author was found: %v", err.Error())
			handler.SendError(context, http.StatusNotFound, "no author was found")
			return
		}

		logger.Errorf("error populating author: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error populating author")
	}

	handler.SendSuccess(context, "list author by id", author)
}
