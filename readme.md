# Well_TCP - Golang 轻量级TCP嵌入式服务器
<p align="center">
    <a href="https://github.com/lvwei25/well_tcp/blob/main/logo.png" target="_blank" style="text-align: center">
        <img src="https://github.com/lvwei25/well_tcp/blob/main/logo.png"  alt="Well_TCP" />
    </a>
</p>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lvwei25/well_tcp)
![GitHub AppVersion](https://img.shields.io/badge/Version-V1.0-blue)

## 介绍

Well_TCP是一款基于Golang的轻量级TCP嵌入式服务器，其为用户内置了路由模块、消息模块、链接模块、分组模块、日志模块，能够满足日常的开发。
<br>


## 使用go get下载依赖

````go
git clone github.com/lvwei25/well_tcp_demo
````

## 关于本项目
本项目是Well_Tcp框架的开发示例，基本开发示例已经写好，用户只需要简单的配置，即可开发项目！
十分的简单方便！

## 配置
````json
{
	"base": {
		"server_name": "Well_TCP",  //服务名称
		"ip_addr": "127.0.0.1",     //服务监听地址
		"port": "6666"              //服务监听端口
	},
	"mysql": {
		"username": "root",         //数据库用户名
		"password": "root",         //数据库密码
		"ip_addr": "127.0.0.1",     //数据库地址
		"port": "3306",             //数据库端口
		"database": "test",         //数据库
		"charset": "utf8"           //数据库编码
	}
}
````

## SQL
本项目数据库方面只用了Mysql，并且是用的GORM库，操作起来是相当方便的，用户只需要在Server启动
时的回调函数对数据库进行连接即可，当然用户可以选择任何地方进行数据库的连接，但是前提是必须在
用户调用s.Run()方法前才可。

> 如何连接? sql.Connect()

> Well 第一个版本发布于 2022 年 3 月 19 日


