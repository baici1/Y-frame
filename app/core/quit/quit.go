package quit

import (
	"Y-frame/app/global/consts"
	"Y-frame/app/global/variable"

	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func init() {
	//  用于系统信号的监听
	go func() {
		//创建一个接受信息号的通道
		c := make(chan os.Signal)
		//	// kill 默认会发送 syscall.SIGTERM 信号
		//	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
		//	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
		//	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM) // 监听可能的退出信号
		received := <-c                                                                           //接收信号管道中的值
		variable.ZapLog.Warn(consts.ProcessKilled, zap.String("信号值", received.String()))          //打印日志
		//(event_manage.CreateEventManageFactory()).FuzzyCall(variable.EventDestroyPrefix)          //销毁事件
		close(c)   //关闭通道
		os.Exit(1) //退出
	}()
}
