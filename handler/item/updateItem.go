package item

import (
	"net/http"
	"strconv"

	"collection.com/config"
	"collection.com/handler"
	"collection.com/models"
	"github.com/gin-gonic/gin"
)

func UpdateItemHandler(context *gin.Context) {
	db := config.GetMySQL()
	logger := config.GetLogger("updateItem")

	request := CreateItemRequest{}
	context.BindJSON(&request)

	id := context.Query("id")

	if id == "" {
		logger.Error("id was not informed")
		handler.SendError(context, http.StatusBadRequest, handler.ErrParamIsRequired("id", "queryParameter").Error())
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

	_, err := db.Exec("UPDATE ITEMS SET "+
		"COLLECTION = ?, TITLE = ?, AUTHOR = ?, PUBLISHER = ?, ITEM_TYPE = ?, ITEM_FORMAT = ?,"+
		"PAGES_NUMBER = ?, EDITION = ?, EDITION_YEAR = ?, BINDING = ?, LANGUAGE = ?, COUNTRY = ?, "+
		"COPIES = ? WHERE ID = ?",
		item.Collection, item.Title, item.Author, item.Publisher, item.ItemType, item.ItemFormat, item.PagesNumber,
		item.Edition, item.EditionYear, item.Binding, item.Language, item.Country, item.Copies, id)

	if err != nil {
		logger.Errorf("error updating item: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		logger.Error("id is not a number")
		handler.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	item.ID = itemId

	handler.SendSuccess(context, "item-updated", item)

}
