package middleware

import (
	"github.com/gin-gonic/gin"
	"web/common"
	"web/components"
)

func Recover(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					//panic(err)
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(appErr)
				return
			}
		}()
		c.Next()
	}
}
