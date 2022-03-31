package routers

import (
	"github.com/lvwei25/well_tcp/log"
	"github.com/lvwei25/well_tcp/net"
)

type ConnectRouter struct {
	net.ConnRouter
}

func (this *ConnectRouter) OnConnect(conn *net.WellConnection) {
	log.NewLoger().Info(conn.IpAddr+":"+conn.Port, "Login")
}
