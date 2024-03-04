package item

import (
	"net/http"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func CreateItemHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("createItem")

	request := CreateItemRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	item := models.Item{
		Collection:  request.Collection, // Collection type
		Title:       request.Title,
		Author:      request.Author,    // Authors type
		Publisher:   request.Publisher, // Publisher type
		ItemType:    request.ItemType,
		ItemFormat:  request.ItemFormat,
		PagesNumber: request.PagesNumber,
		Edition:     request.Edition,
		EditionYear: request.EditionYear,
		Binding:     request.Binding,
		Language:    request.Language, // Language type
		Country:     request.Country,  // Country type
		Copies:      request.Copies,
	}

	result, err := db.Exec("INSERT INTO ITEMS"+
		"(COLLECTION, TITLE, AUTHOR, PUBLISHER, ITEM_TYPE, ITEM_FORMAT,"+
		"PAGES_NUMBER, EDITION, EDITION_YEAR, BINDING, LANGUAGE, COUNTRY, COPIES) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		item.Collection, item.Title, item.Author, item.Publisher, item.ItemType, item.ItemFormat, item.PagesNumber,
		item.Edition, item.EditionYear, item.Binding, item.Language, item.Country, item.Copies)

	if err != nil {
		logger.Errorf("error creating item: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Errorf("error creating item: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	item.ID = id

	handler.SendSuccess(context, "item created", item)

}
