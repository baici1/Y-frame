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
	//从context获取参数
	UserName := c.GetString("form" + "user_name")
	pass := c.GetString("form" + "pass")
	//创建数据库操作对象
	userModelFact := model.CreateUserFactory("")
	//进入model层
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
