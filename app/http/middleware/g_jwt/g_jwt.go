package g_jwt

import (
	"Y-frame/app/global/g_errors"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims 自定义jwt的声明字段 （注：不能存储敏感信息）
type CustomClaims struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"user_name"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}

//CreateGJWT
/* @Description: 创建一个生成jwt签名的结构体工厂
 * @param signKey
 * @return *JwtSign
 */
func CreateGJWT(signKey string) *JwtSign {
	if len(signKey) == 0 {
		signKey = "Y-frame"
	}
	return &JwtSign{
		SigningKey: []byte(signKey),
	}
}

// JwtSign 定义一个 JWT验签 结构体
type JwtSign struct {
	SigningKey []byte
}

//CreateJWT
/* @Description: 生成token字符串
 * @receiver j
 * @param claims 负载
 * @return string token字符串
 * @return error
 */
func (j *JwtSign) CreateJWT(claims CustomClaims) (string, error) {
	// 使用指定的签名方法创建签名对象 (选择加密技术)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token  （添加签名值，生成完整的token）
	return token.SignedString(j.SigningKey)
}

//ParseJWT
/* @Description: 解析token
 * @receiver j
 * @param tokenString  token字符串
 * @return *CustomClaims 返回token的负载
 * @return error
 */
func (j *JwtSign) ParseJWT(tokenString string) (*CustomClaims, error) {
	//解析token
	/*
		三个参数：
		tokenString string token字符串
		claims  存储解析后的jwt结构体
		keyFunc  返回生成jwt设置的secret 起到验证签名作用
	*/
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	/*
		token字段
			Raw       string                 // 原始令牌。解析令牌时填充
			Method    SigningMethod          // 使用或将使用的签名方法
			Header    map[string]interface{} // 令牌的第一段
			Claims    Claims                 // 令牌的第二部分
			Signature string                 // 令牌的第三部分。解析令牌时填充
			Valid     bool                   // 令牌有效吗？解析/验证令牌时填充
	*/
	//token无效
	if token == nil {
		return nil, errors.New(g_errors.ErrorsTokenInvalid)
	}
	if err != nil {
		//校验失败，分情况 。超时要特殊处理 通过与运算来进行与token的错误比对
		if ve, ok := err.(*jwt.ValidationError); ok {
			//格式不正确
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(g_errors.ErrorsTokenMalFormed)
				//token未激活
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(g_errors.ErrorsTokenNotActiveYet)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //如果token只是过期，那么我们认为他是可用token 续期就可
				token.Valid = true
			} else {
				return nil, errors.New(g_errors.ErrorsTokenInvalid)
			}
		}
	}
	//获取token的负载
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New(g_errors.ErrorsTokenInvalid)
	}
}

//RefreshJWT
/* @Description: 更新token
 * @receiver j
 * @param tokenString  token字符串
 * @param extraAddSeconds 额外增加的时间
 * @return string 返回新的token
 * @return error 报错
 */
func (j *JwtSign) RefreshJWT(tokenString string, extraAddSeconds int64) (string, error) {
	//解析token 看是否是过期情况
	if CustomClaims, err := j.ParseJWT(tokenString); err != nil {
		return "", err
	} else {
		//设置过期时间
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		//重新创建新的token
		return j.CreateJWT(*CustomClaims)
	}
}
