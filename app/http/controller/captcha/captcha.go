package captcha

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	"Y-frame/app/utils/response"
	"bytes"
	"net/http"
	"path"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

//设置获取请求的字段key
const (
	CaptchaIdKey    string = "captcha_id"
	CaptchaValueKey string = "captcha_value"
	Reload          string = "reload"
)

//获取配置文件中的设置
var (
	Width  int    = variable.ConfigYml.GetInt("Captcha.StdWidth")
	Height int    = variable.ConfigYml.GetInt("Captcha.StdHeight")
	Lang   string = variable.ConfigYml.GetString("Captcha.Lang")
)

type Captcha struct {
	Id      string `json:"id,omitempty"`      //验证码ID
	ImgUrl  string `json:"img_url,omitempty"` //验证码图像地址
	Refresh string `json:"refresh,omitempty"` //重新获取
	Verify  string `json:"verify,omitempty"`  //验证
}

func init() {
	//对存储器进行自定义设置
	//设置一次清理过期验证码的数量
	collectNum := variable.ConfigYml.GetInt("Captcha.CollectNum")
	//过期时间
	expiration := variable.ConfigYml.GetDuration("Captcha.Expiration")
	// 返回一个新的标准内存存储器
	s := captcha.NewMemoryStore(collectNum, expiration)
	//设置一个新的存储器
	captcha.SetCustomStore(s)
}

//GenerateId
/* @Description: 生成验证码Id
 * @receiver c
 * @param ctx
 */
func (c *Captcha) GenerateId(ctx *gin.Context) {
	//获取验证码验证数字长度
	var length = variable.ConfigYml.GetInt("Captcha.Length")
	//自定义验证码数字长度
	captchaId := captcha.NewLen(length)
	//提供相关信息
	c.Id = captchaId
	c.ImgUrl = captchaId + ".png"
	c.Refresh = c.ImgUrl + "?reload=1"
	c.Verify = captchaId + "/这里替换为正确的验证码进行验证"
	response.Success(ctx, "验证码信息", c)
}

//GetImg
/* @Description: 生成验证码的图片
 * @receiver c
 * @param ctx
 */
func (c *Captcha) GetImg(ctx *gin.Context) {
	//通过param获取验证码Id
	captchaId := ctx.Param(CaptchaIdKey)
	//根据路径的最后一个斜杠， 将路径划分为目录部分和文件名部分。 例 http://localhost:20201/captcha/  JUpvESBqAcPxvSVz5Cy9.png
	_, file := path.Split(ctx.Request.URL.Path)
	//获取扩展名 .png
	ext := path.Ext(file)
	//获取路径上的Id
	id := file[:len(file)-len(ext)]
	//验证参数是否获取正常
	if ext == "" || captchaId == "" {
		response.Fail(ctx, consts.CaptchaGetParamsInvalidCode, consts.CaptchaGetParamsInvalidMsg)
		return
	}
	//当获取到 Reload 就重新加载一个
	if ctx.Query(Reload) != "" {
		//重载为给定的验证码生成并记住新的数字。
		captcha.Reload(id)
	}
	//设置http协议
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Expires", "0")
	var vBytes bytes.Buffer
	if ext == ".png" {
		ctx.Header("Content-Type", "image/png")
		//写入图片的二进制信息
		_ = captcha.WriteImage(&vBytes, id, Width, Height)
		//读取文件内容并输出的方法
		http.ServeContent(ctx.Writer, ctx.Request, id+ext, time.Time{}, bytes.NewReader(vBytes.Bytes()))
	}
}

//CheckCode
/* @Description: 校验验证码传来的数字
 * @receiver c
 * @param ctx
 */
func (c *Captcha) CheckCode(ctx *gin.Context) {
	//获取验证码ID 与验证码的值
	captchaId := ctx.Param(CaptchaIdKey)
	value := ctx.Param(CaptchaValueKey)
	//去存储器中进行验证
	if captcha.VerifyString(captchaId, value) {
		response.Success(ctx, consts.CaptchaCheckParamsOk)
	} else {
		response.Fail(ctx, consts.CaptchaCheckParamsInvalidCode, consts.CaptchaCheckParamsInvalidMsg)
	}
}

//GetAudio
/* @Description: 创建audio的验证码
 * @receiver c
 * @param ctx
 */
func (c *Captcha) GetAudio(ctx *gin.Context) {
	//通过param获取验证码Id
	captchaId := ctx.Param(CaptchaIdKey)
	//根据路径的最后一个斜杠， 将路径划分为目录部分和文件名部分。 例 http://localhost:20201/captcha/  JUpvESBqAcPxvSVz5Cy9.wav
	_, file := path.Split(ctx.Request.URL.Path)
	//获取扩展名 .wav
	ext := path.Ext(file)
	//获取路径上的Id
	id := file[:len(file)-len(ext)]
	//验证参数是否获取正常
	if ext == "" || captchaId == "" {
		response.Fail(ctx, consts.CaptchaGetParamsInvalidCode, consts.CaptchaGetParamsInvalidMsg)
		return
	}
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	//当获取到 Reload 就重新加载一个
	if ctx.Query(Reload) != "" {
		//重载为给定的验证码生成并记住新的数字。
		captcha.Reload(id)
	}
	var vBytes bytes.Buffer
	if ext == ".wav" {
		ctx.Header("Content-Type", "audio/wav")
		//写入audio的二进制信息
		_ = captcha.WriteAudio(&vBytes, id, Lang)
		//读取文件内容并输出的方法
		http.ServeContent(ctx.Writer, ctx.Request, id+ext, time.Time{}, bytes.NewReader(vBytes.Bytes()))
	}
}
