package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string      `json:"error"`
	More  interface{} `json:"more"`
}

func HandleErrorInHandler(c *gin.Context, err error, more interface{}) {
	c.Error(err)
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error: err.Error(),
		More:  more,
	})
}

func OrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
