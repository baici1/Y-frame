package token

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	"Y-frame/app/http/middleware/g_jwt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CreateUserToken
/* @Description: 生成一个userToken对象（管理token）
 * @return *userToken
 */
func CreateUserToken() *userToken {
	return &userToken{
		userJWT: g_jwt.CreateGJWT(variable.ConfigYml.GetString("Token.JwtTokenSignKey")),
	}
}

type userToken struct {
	userJWT *g_jwt.JwtSign
}

//GenerateToken
/* @Description: 生成token
 * @receiver u
 * @param username 用户名
 * @param phone 电话
 * @param userid 用户ID
 * @param expireAt 有效期
 * @return token
 * @return err
 */
func (u *userToken) GenerateToken(username, phone string, userid, expireAt int64) (token string, err error) {
	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := g_jwt.CustomClaims{
		UserId: userid,
		Name:   username,
		Phone:  phone,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
		},
	}
	//根据负载，生成token
	return u.userJWT.CreateJWT(customClaims)
}

//ParseToken
/* @Description:解析token
 * @receiver u
 * @param tokenString token字符串
 * @return g_jwt.CustomClaims 负载结构体
 * @return error
 */
func (u *userToken) ParseToken(tokenString string) (g_jwt.CustomClaims, error) {
	if customClaims, err := u.userJWT.ParseJWT(tokenString); err != nil {
		return g_jwt.CustomClaims{}, nil
	} else {
		return *customClaims, nil
	}
}

//RefreshToken
/* @Description: 更新token
 * @receiver u
 * @param oldToken 旧的token
 * @return newToken 新的token
 * @return flag
 */
func (u *userToken) RefreshToken(oldToken string) (newToken string, flag bool) {
	extraAddSeconds := variable.ConfigYml.GetInt64("Token.JwtTokenRefreshExpireAt")
	if newToken, err := u.userJWT.RefreshJWT(oldToken, extraAddSeconds); err != nil {
		return "", false
	} else {
		return newToken, true
	}

}

//IsEffect
/* @Description: 判断token是否失效
 * @receiver u
 * @param token
 * @return *g_jwt.CustomClaims
 * @return int
 */
func (u *userToken) IsEffect(token string) (*g_jwt.CustomClaims, int) {
	if customClaims, err := u.userJWT.ParseJWT(token); err != nil {
		return nil, consts.JwtTokenInvalid
	} else {
		//是否在有效期
		if time.Now().Unix()-customClaims.ExpiresAt < 0 {
			//还在有效期内
			return customClaims, consts.JwtTokenOK
		} else {
			return customClaims, consts.JwtTokenExpired
		}
	}
}
