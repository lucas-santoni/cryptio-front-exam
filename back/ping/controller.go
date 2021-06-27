package ping

import (
	"back/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

const pong = "pong\n"

func GET(_api *api.API) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, pong)
	}
}
