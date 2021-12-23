package snow_flake

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"
	"Y-frame/app/utils/snow_flake/snow_flake_interf"
	"sync"
	"time"
)

//创建一个雪花算法生成器工程
func CreateSnowFlakeFactory() snow_flake_interf.InterfaceSnowFlake {
	return &snowflake{
		timestamp: 0,
		machineId: variable.Configs.SnowFlake.SnowFlakeMachineId,
		sequence:  0,
	}
}

type snowflake struct {
	sync.Mutex       //添加互斥锁，确保并发安全性
	timestamp  int64 //记录上一次生成ID的时间戳
	machineId  int64 //机器号
	sequence   int64 //当前毫秒已经生成的ID序列号(从0 开始累加) 1毫秒内最多生成4096个ID
}

// 生成分布式ID
func (s *snowflake) GetId() int64 {
	//加锁
	s.Lock()
	defer func() {
		s.Unlock()
	}()
	//获取当前时间的时间戳（毫秒）
	now := time.Now().UnixNano() / 1e6
	//当前时间与工作节点上一次生成ID的时间
	if s.timestamp == now {
		//相当于sequence++，也检测是否溢出
		s.sequence = (s.sequence + 1) & consts.SequenceMask
		//毫秒内序列溢出
		if s.sequence == 0 {
			//阻塞到下一个毫秒,获得新的时间戳
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	//记录上一次生成ID的时间戳
	s.timestamp = now
	//移位并通过或运算拼到一起组成64位的ID
	r := (now-consts.StartTimeStamp)<<consts.TimestampShift | (s.machineId << consts.MachineIdShift) | (s.sequence)
	return r
}
