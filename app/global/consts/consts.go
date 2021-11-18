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
	ProcessKilled = "系统退出，信号是："

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

	//response状态码
	CurdStatusOkCode  int    = 200 //请求成功
	CurdStatusOkMsg   string = "请求成功"
	CurdLoginFailCode int    = -400100
	CurdLoginFailMsg  string = "登录失败"
	//表单验证
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"
)
