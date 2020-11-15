package server

import (
	"net/http"
	"yals/db"
	"yals/utils"

	"github.com/gin-gonic/gin"
)

// HandleIndex - handles index.
func HandleIndex(c *gin.Context) {
	c.String(http.StatusOK, "<h1>Hello world!</h2>")
}

// HandlePostRequest - handler for adding item to database.
func HandlePostRequest(c *gin.Context) {
	var postJSON = make(map[string]string)
	c.BindJSON(&postJSON)
	URL := postJSON["url"]
	if utils.IsURL(URL) == true {
		item, err := db.GetOrCreateUsingURL(URL)
		if err != nil {
			errorResponse := map[string]string{"error": err.Error()}
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
		c.JSON(http.StatusOK, item)
	} else {
		errorResponse := map[string]string{"error": "Invalid URL"}
		c.JSON(http.StatusOK, errorResponse)
	}
}

// HandleGetRequest - handler for getting item from database.
func HandleGetRequest(c *gin.Context) {
	Shortcut := c.Param("Shortcut")
	item, err := db.GetByShortcut(Shortcut)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, item.URL)
	}
}
