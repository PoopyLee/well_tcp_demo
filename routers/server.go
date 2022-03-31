package routers

import (
	"github.com/lvwei25/well_tcp/log"
	"well_tcp/net"
)

type ServerRouter struct {
	net.WellServerRouter
}

func (this *ServerRouter) OnStart() {
	log.NewLoger().Info("This is Onstart")
	//sql.RedisConnect()
	//sql.Connect()
	//sql.Db.AutoMigrate(&models.Default{})
}
