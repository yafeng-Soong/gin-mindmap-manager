package middleware

import (
	"paper-manager/model/common/response"
	"paper-manager/model/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if length := len(c.Errors); length > 0 {
			e := c.Errors[length-1]
			err := e.Err
			if err != nil {
				if myErr, ok := err.(*errors.MyError); ok {
					response.FailWithError(myErr, c)
				} else {
					response.ServerError(err.Error(), c)
				}
			}
		}
	}
}
