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
	//s:=models.PrintAll()
	//for _,v:=range s{
	//	conn.WriteString(fmt.Sprintf("%d->%s",v.ID,v.Name)+"\r\n")
	//}
}
