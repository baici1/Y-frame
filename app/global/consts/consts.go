package consts

//定义一些常量 一般是具有错误代码+错误说明组成，一般用于接口返回 有一些信息提示

const (

	//配置文件发生变化
	YamlConfigChange = "配置文件发生了变化"
	//雪花算法相关参数
	StartTimeStamp = int64(1483228800000) //开始时间截 (2017-01-01)
	MachineIdBits  = uint(10)             //机器id所占的位数
	SequenceBits   = uint(12)             //序列所占的位数
	//MachineIdMax   = int64(-1 ^ (-1 << MachineIdBits)) //支持的最大机器id数量
	SequenceMask   = int64(-1 ^ (-1 << SequenceBits)) //
	MachineIdShift = SequenceBits                     //机器id左移位数
	TimestampShift = SequenceBits + MachineIdBits     //时间戳左移位数

	//系统参数
	ProcessKilled                  = "系统退出，信号是："
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误, "

	//token状态码
	JwtTokenOK              int    = 200100  //token有效
	JwtTokenInvalid         int    = -400100 //无效的token
	JwtTokenInvalidMsg      string = "无效的oken"
	JwtTokenExpired         int    = -400101 //过期的token
	JwtTokenFormatErrCode   int    = -400102 //提交的 token 格式错误
	JwtTokenFormatErrMsg    string = "token格式错误"
	JwtTokenMustValidMsg    string = "token为必填项,请在请求header部分提交!" //提交的 token 格式错误
	JwtTokenRefreshFailCode int    = -400200
	JwtTokenRefreshFailMsg  string = "token刷新失败，请重新登录"

	//表单验证
	ValidatorPrefix              string = "Form_Validator_" //表单验证前缀
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"
	ValidatorParamsToJSONFail    string = "验证器参数 json 化失败"
	//验证码
	CaptchaGetParamsInvalidMsg    string = "获取验证码：提交的验证码参数无效,请检查验证码ID以及文件名后缀是否完整"
	CaptchaGetParamsInvalidCode   int    = -400350
	CaptchaCheckParamsInvalidMsg  string = "校验验证码：提交的参数无效或者验证码已失效，请检查 【验证码ID、验证码值】或者改配置文件中的过期时间。"
	CaptchaCheckParamsInvalidCode int    = -400351
	CaptchaCheckParamsOk          string = "验证通过"
	CaptchaCheckParamsFailMsg     string = "验证码错误"
	CaptchaCheckParamsFailCode    int    = -400200

	//用户状态码
	UserLoginFailCode    int    = -400100
	UserLoginFailMsg     string = "登录失败"
	UserRegisterFailCode int    = -400200
	UserRegisterFailMsg  string = "注册失败"
	//CURD常用状态码
	CurdStatusOkCode   int    = 200
	CurdStatusOkMsg    string = "Success"
	CurdSelectFailCode int    = -400201 //关于查询的
	CurdSelectFailMsg  string = "查询失败"
)
