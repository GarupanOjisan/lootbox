package main

import (
	"github.com/gin-gonic/gin"

	"github.com/garupanojisan/lootbox/app/admin-api/lootbox"
)

func main() {
	r := gin.Default()

	g := r.Group("/lootbox")
	g.GET("", lootbox.Index)

	r.Run(":8081")
}
