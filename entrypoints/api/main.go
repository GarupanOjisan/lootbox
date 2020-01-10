package main

import (
	"github.com/gin-gonic/gin"

	"github.com/garupanojisan/omikuji-api/app/api/lootbox"
)

func main() {
	r := gin.Default()

	g := r.Group("/lootbox")
	g.POST("", lootbox.Post)

	r.Run(":8080")
}
