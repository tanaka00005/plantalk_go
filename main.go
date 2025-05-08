package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Hello, World!")

	const URL = "https://search.rakuten.co.jp/search/mall/keyboard/";


	r := gin.Default()

	r.GET(URL,func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"hello world",
		})
	})
	
}




