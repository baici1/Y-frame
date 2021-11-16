package web

import (
	"Y-frame/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
}

//用户登录
func (u *Users) Login(c *gin.Context) {
	UserName := c.GetString("form" + "user_name")
	pass := c.GetString("form" + "pass")
	userModelFact := model.CreateUserFactory("")
	userModel := userModelFact.Login(UserName, pass)
	if userModel != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": userModel,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "失败",
	})
}
