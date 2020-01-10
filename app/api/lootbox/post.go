package lootbox

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/garupanojisan/lootbox/app/usecase/lootbox"
)

func Post(c *gin.Context) {
	u := &lootbox.UseCase{}

	num, err := strconv.Atoi(c.PostForm("num"))
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	items, err := u.Pick(num)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}
