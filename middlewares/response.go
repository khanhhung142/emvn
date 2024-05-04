package middlewares

import (
	"emvn/consts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if err, ok := c.Get(consts.GinErrorKey); ok {
			err := err.(consts.CustomError)
			if additionalErr, ok := c.Get(consts.GinDetailErrorKey); ok {
				err.Message = err.Message + " | " + additionalErr.(error).Error()
			}
			c.JSON(err.HttpStatus, err.Detail())
			c.Abort()
			return
		}

		if data, ok := c.Get(consts.GinResponseKey); ok {
			c.JSON(http.StatusOK, Response{
				Code:    0,
				Message: "Success",
				Data:    data,
			})
		}
	}
}
