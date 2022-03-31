package sql

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/lvwei25/well_tcp/log"
	"github.com/lvwei25/well_tcp_demo/config"
)

var Db *gorm.DB
var Rdb *redis.Client

func MysqlConnect() error {
	conf := config.ReturnConfig()
	dns := conf.Mysql.Username + ":" + conf.Mysql.Password + "@tcp(" +
		conf.Mysql.IpAddr + ":" + conf.Mysql.Port + ")/" + conf.Mysql.Database + "?charset=" + conf.Mysql.Charset + "&parseTime=True&loc=Local"
	log.NewLoger().Info(dns)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.NewLoger().Error(err)
		return err
	}
	log.NewLoger().Info(db, conf.Mysql.Database, "DataBase Connect Success")
	Db = db
	return nil
}

func RedisConnect() error {
	conf := config.ReturnConfig()
	// 通过 redis.NewClient 函数即可创建一个 redis 客户端, 这个方法接收一个 redis.Options 对象参数, 通过这个参数, 我们可以配置 redis 相关的属性, 例如 redis 服务器地址, 数据库名, 数据库密码等。
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.IpAddr + ":" + conf.Redis.Port,
		Password: conf.Redis.Password, // no password set
		DB:       0,                   // use default DB
	})

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	res, err := rdb.Ping().Result()
	log.NewLoger().Info(res)
	if err != nil {
		return err
	}
	Rdb = rdb
	return nil
}

func MysqlStop() {
	Db.Close()
}

func RedisStop() {
	Rdb.Close()
}
