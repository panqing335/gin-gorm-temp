package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetID(c *gin.Context) int64 {
	value := c.GetString("Id")
	i, _ := strconv.ParseInt(value, 10, 64)

	return i
}
