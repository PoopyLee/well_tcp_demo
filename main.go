package main

import (
	"github.com/lvwei25/well_tcp/net"
	config "github.com/lvwei25/well_tcp_demo/config"
	"github.com/lvwei25/well_tcp_demo/routers"
)

func main() {
	con := config.ReturnConfig()
	s := net.NewServerHandle(con.Base.ServerName, con.Base.IpAddr, con.Base.Port)
	s.AddServerRouter(&routers.ServerRouter{})
	s.AddConnRouter(&routers.ConnectRouter{})
	s.Run()
}
