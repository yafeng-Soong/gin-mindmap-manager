package middleware

import (
	"github.com/yafeng-Soong/gin-mindmap-manager/model/common/response"
	"github.com/yafeng-Soong/gin-mindmap-manager/model/errors"

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
