package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	models "go_web_app_beego/models"
	_ "go_web_app_beego/routers"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/products.db")
	orm.RegisterModel(new(models.Product))
}

func main() {
	beego.Run()
}
