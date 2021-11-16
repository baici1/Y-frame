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
)
