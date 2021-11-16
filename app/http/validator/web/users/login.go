package users

import (
	"Y-frame/app/http/controller/web"
	"Y-frame/app/http/validator/core/data_transfer"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用户登录的参数校验
type Login struct {
	UserName string `form:"user_name" json:"user_name"  binding:"required,min=1"` // 必填、对于文本,表示它的长度>=1
	Pass     string `form:"pass" json:"pass" binding:"required,min=6,max=20"`     //  密码为 必填，长度>=6

}

//验证器
func (l Login) CheckParams(c *gin.Context) {
	//获取参数值，进行初步的验证规则
	if err := c.ShouldBind(&l); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "失败1",
		})
		c.Abort()
		return
	}
	//将参数值值绑定到context上下文中
	extraAddBindDataContext := data_transfer.DataAddContext(l, "form", c)
	if extraAddBindDataContext != nil {
		(&web.Users{}).Login(extraAddBindDataContext)
	}
}
