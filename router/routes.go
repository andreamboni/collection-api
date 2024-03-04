package router

import (
	"collection.com/author"
	"collection.com/collection"
	"collection.com/country"
	"collection.com/item"
	"collection.com/language"
	"collection.com/publisher"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/collection-ws/v1"
	v1 := router.Group(basePath)
	{
		// item
		v1.GET("/items", item.ListItemsHandler)
		v1.GET("/itemById", item.ListItemByIdHandler)
		v1.GET("/deletedItems", item.ListDeletedItems)
		v1.POST("/createItem", item.CreateItemHandler)
		v1.PUT("/updateItem", item.UpdateItemHandler)
		v1.PUT("/recoverItem", item.RecoverItemHandler)
		v1.DELETE("/deleteItem", item.DeleteItemHandler)

		// language
		v1.GET("/languages", language.ListLanguageHandler)
		v1.GET("/languagesById", language.ListLanguageByIdHandler)
		v1.GET("/deletedLanguages", language.ListDeletedLanguages)
		v1.POST("/createLanguage", language.CreateLanguageHandler)
		v1.PUT("/updateLanguage", language.UpdateLanguageHandler)
		v1.PUT("/recoverLanguage", language.RecoverLanguageHandler)
		v1.DELETE("/deleteLanguage", language.DeleteLanguageHandler)

		// author
		v1.GET("/authors", author.ListAuthorsHandler)
		v1.GET("/authorById", author.ListAuthorByIdHandler)
		v1.GET("/deletedAuthors", author.ListDeletedAuthors)
		v1.POST("/createAuthor", author.CreateAuthorHandler)
		v1.PUT("/updateAuthor", author.UpdateAuthorHandler)
		v1.PUT("/recoverAuthor", author.RecoverAuthorHandler)
		v1.DELETE("/deleteAuthor", author.DeleteAuthorHandler)

		// collection
		v1.GET("/collections", collection.ListCollectionsHandler)
		v1.GET("/collectionById", collection.ListCollectionByIdHandler)
		v1.GET("/deletedCollections", collection.ListDeletedCollection)
		v1.POST("/createCollection", collection.CreateCollectionHandler)
		v1.PUT("/updateCollection", collection.UpdateCollectionHandler)
		v1.PUT("/recoverCollection", collection.RecoverCollectionHandler)
		v1.DELETE("/deleteCollection", collection.DeleteCollectionHandler)

		// country
		v1.GET("/countries", country.ListCountriesHandler)
		v1.GET("/countryById", country.ListCountryByIdHandler)
		v1.GET("/deletedCountries", country.ListDeletedCountries)
		v1.POST("/createCountry", country.CreateCountryHandler)
		v1.PUT("/updateCountry", country.UpdateCountryHandler)
		v1.PUT("/recoverCountry", country.RecoverCountryHandler)
		v1.DELETE("/deleteCountry", country.DeleteCountryHandler)

		// publisher
		v1.GET("/publishers", publisher.ListPublishersHandler)
		v1.GET("/publisherById", publisher.ListPublisherByIdHandler)
		v1.GET("/deletedPublishers", publisher.ListDeletedPublishers)
		v1.POST("/createPublisher", publisher.CreatePublisherHandler)
		v1.PUT("/updatePublisher", publisher.UpdatePublisherHandler)
		v1.PUT("/recoverPublisher", publisher.RecoverPublisherHandler)
		v1.DELETE("/deletePublisher", publisher.DeletePublisherHandler)
	}
}
