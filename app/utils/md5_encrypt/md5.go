package md5_encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

//MD5
/* @Description: MD5加密
 * @param params
 * @return string
 */
func MD5(params string) string {
	//生成一个对象
	md5Ctx := md5.New()
	//写入字符流
	md5Ctx.Write([]byte(params))
	//返回字符串
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

//Base64Md5
/* @Description: 先进行base64加密，然后MD5加密
 * @param params
 * @return string
 */
func Base64Md5(params string) string {
	//先进行base64编码，然后在加密
	return MD5(base64.StdEncoding.EncodeToString([]byte(params)))
}
