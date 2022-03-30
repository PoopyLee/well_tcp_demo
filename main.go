package main

import (
	"github.com/lvwei25/well_tcp/net"
	"well_tcp/log"
	config "well_tcp_demo/config"
	"well_tcp_demo/routers"
)

func main() {
	con := config.ReturnConfig()
	log.NewLoger().SetLogFile("logs.log")
	s := net.NewServerHandle(con.Base.ServerName, con.Base.IpAddr, con.Base.Port)
	s.AddServerRouter(&routers.ServerRouter{})
	s.AddConnRouter(&routers.ConnectRouter{})
	s.Run()
}
