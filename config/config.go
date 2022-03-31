package config

import (
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvwei25/well_tcp/log"
	"io/ioutil"
	"os"
)

var conf Config

type Config struct {
	Base  baseConfig  `json:"base"`
	Mysql mysqlConfig `json:"mysql"`
	Redis redisConfig `json:"redis"`
}

type baseConfig struct {
	ServerName string `json:"server_name"` //服务名称
	IpAddr     string `json:"ip_addr"`     //监听地址
	Port       string `json:"port"`        //监听端口
}

type mysqlConfig struct {
	Username string `json:"username"` //数据库账号
	Password string `json:"password"` //数据库密码
	IpAddr   string `json:"ip_addr"`  //数据库IP地址
	Port     string `json:"port"`     //数据库端口
	Database string `json:"database"` //连接的数据库
	Charset  string `json:"charset"`  //编码
}

type redisConfig struct {
	Password string `json:"password"` //数据库密码
	IpAddr   string `json:"ip_addr"`  //数据库IP地址
	Port     string `json:"port"`     //数据库端口
}

func init() {
	//os.Mkdir("conf", 755)
	//file, err := os.Open("conf/config.json")
	//if err != nil {
	//	os.Mkdir("conf", 755)
	//	file, _ = os.Create("conf/config.json")
	//	log.NewLoger().Error(err)
	//
	//}
	//defer file.Close()

	data, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		os.Mkdir("conf", 755)
		os.Create("conf/config.json")
		log.NewLoger().Error(err)
		return
	}
	if len(data) > 0 {
		err := json.Unmarshal(data, &conf)
		if err != nil {
			log.NewLoger().Error(err)
		}
	} else {
		config := new(Config)
		//=======================
		config.Base.IpAddr = "127.0.0.1"
		config.Base.Port = "6666"
		config.Base.ServerName = "Well_TCP"
		//=======================
		config.Mysql.IpAddr = "127.0.0.1"
		config.Mysql.Port = "3306"
		config.Mysql.Username = "root"
		config.Mysql.Password = "root"
		config.Mysql.Database = "test"
		config.Mysql.Charset = "utf8"
		//=======================
		config.Redis.IpAddr = "localhost"
		config.Redis.Port = "6379"
		config.Redis.Password = "root"
		config_s, _ := json.MarshalIndent(config, "", "\t")
		err := ioutil.WriteFile("conf/config.json", config_s, 0777)
		if err != nil {
			log.NewLoger().Error(err)
		}
		return
	}
}

func ReturnConfig() Config {
	return conf
}
