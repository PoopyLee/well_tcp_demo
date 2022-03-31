package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvwei25/well_tcp/log"
	"github.com/lvwei25/well_tcp_demo/sql"
	"time"
)

type Default struct {
	gorm.Model
	Name string
	Time time.Time
}

func PrintAll() []Default {
	tests := make([]Default, 1)
	sql.Db.Find(&tests)
	for k, v := range tests {
		log.NewLoger().Info(k, v)
	}
	return tests
}
