package portfolio

import (
	"back/api"
	"back/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETHistory(api *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		entries, err := history()
		if err != nil {
			utils.HandleErrorInHandler(c, err, nil)
			return
		}

		c.JSON(http.StatusOK, entries)
	}
}

func GETAssets(api *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		entries, err := topAssets()
		if err != nil {
			utils.HandleErrorInHandler(c, err, nil)
			return
		}

		c.JSON(http.StatusOK, entries)
	}
}
