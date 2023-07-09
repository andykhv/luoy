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
		value, ok := cache.Get(key)

		if ok {
			context.IndentedJSON(http.StatusFound, value)
		} else {
			context.String(http.StatusNotFound, fmt.Sprintf("Key {%v} not found", key))
		}
	})

	router.PUT("/cache/:key", func(context *gin.Context) {
		key := context.Param("key")
		var value interface{}
		context.BindJSON(&value)

		ok := cache.Insert(key, value)

		if ok {
			context.IndentedJSON(http.StatusOK, value)
		} else {
			context.String(http.StatusInternalServerError, "internal error, sorry!")
		}
	})

	router.Run(address)
}
