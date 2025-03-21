package Utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ValidateAndExtractIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			c.Abort() // Прерываем запрос
			return
		}
		c.Set("ID", int64(id)) // Сохраняем ID в контексте
		c.Next()               // Продолжаем выполнение следующего middleware или хендлера
	}
}
