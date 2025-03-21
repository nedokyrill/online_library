package Utils

import "github.com/gin-gonic/gin"

func GetIDFromContext(c *gin.Context) int64 {
	ID, isExist := c.Get("ID")
	if !isExist {
		return 0
	}
	return ID.(int64)
}
