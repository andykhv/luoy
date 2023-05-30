package luoy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLuoy(address string) {
	cache := NewCache()
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "hello!")
	})

	router.GET("/cache/:key", func(context *gin.Context) {
		key := context.Param("key")
		value, ok := cache.data[key]

		if ok {
			context.IndentedJSON(http.StatusFound, value)
		} else {
			context.String(http.StatusNotFound, fmt.Sprintf("Key {%v} not found", key))
		}
	})

	router.Run(address)
}
