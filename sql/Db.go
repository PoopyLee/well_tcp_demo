package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/lvwei25/well_tcp/log"
	"well_tcp_demo/config"
)

var Db *gorm.DB

func Connect() {
	conf := config.ReturnConfig()
	dns := conf.Mysql.Username + ":" + conf.Mysql.Password + "@tcp(" +
		conf.Mysql.IpAddr + ":" + conf.Mysql.Port + ")/" + conf.Mysql.Database + "?charset=" + conf.Mysql.Charset + "&parseTime=True&loc=Local"
	log.NewLoger().Info(dns)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.NewLoger().Error(err)
		panic(err)
	}
	log.NewLoger().Info(db, conf.Mysql.Database, "DataBase Connect Success")
	Db = db
	//Db.Debug()
}

func Stop() {
	Db.Close()
}
