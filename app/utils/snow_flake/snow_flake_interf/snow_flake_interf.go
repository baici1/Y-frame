package snow_flake_interf

//雪花算法对外的接口
type InterfaceSnowFlake interface {
	GetId() int64
}
