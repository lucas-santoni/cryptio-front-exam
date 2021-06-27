package portfolio

import (
	"back/api"
	"back/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GETHistory godoc
// @Summary Retrieve the portfolio history of an user
// @Description 2 months worth of data are returned at most, which is around 60 entries.
// @Description Valuation is expressed in USD.
// @Produce json
// @Success 200 {object} []PortfolioHistoryEntry
// @Failure 500 {object} utils.ErrorResponse
// @Router /portfolio/history [get]
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

// GETHistory godoc
// @Summary Retrieve the top assets of the portfolio of an user
// @Description 9 assets are returned. The 10th entry aggregates all the other assets.
// @Description Volumes and prices are expressed in USD.
// @Produce json
// @Success 200 {object} []PortfolioTopAssetEntry
// @Failure 500 {object} utils.ErrorResponse
// @Router /portfolio/top-assets [get]
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
