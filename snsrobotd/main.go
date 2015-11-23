package main

import (
	_ "github.com/quchunguang/snsrobot-go/snsrobotd/docs"
	_ "github.com/quchunguang/snsrobot-go/snsrobotd/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:123456@172.17.0.2:5432/snsrobot?sslmode=disable")
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

