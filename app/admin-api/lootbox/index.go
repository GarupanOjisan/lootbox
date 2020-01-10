package lootbox

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/garupanojisan/omikuji-api/app/usecase/lootbox"
)

func Index(c *gin.Context) {
	u := lootbox.UseCase{}

	userId := ""
	id := c.Query("id")
	items, err := u.GetGroup(userId, id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}
