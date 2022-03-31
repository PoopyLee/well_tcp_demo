package routers

import (
	"github.com/lvwei25/well_tcp/log"
	"github.com/lvwei25/well_tcp/net"
)

type ConnectRouter struct {
	net.ConnRouter
}

//TCP握手三次成功之后 第一次执行的回调
func (this *ConnectRouter) OnConnect(con *net.WellConnection) {
	log.NewLoger().Info(con.IpAddr+":"+con.Port, "Login")
}

//发送信息的回调
func (this *ConnectRouter) OnMessage(request *net.WellRequest) {
	log.NewLoger().Info(request.IpAddr+":"+request.Port+"say:", request.GetData(), "\r\n")
}

//连接断开的回调
func (this *ConnectRouter) OnClose(con *net.WellConnection) {
	log.NewLoger().Info(con.IpAddr+":"+con.Port, "Logout")
}
