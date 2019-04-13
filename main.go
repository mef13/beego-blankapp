package main

import (
	"beego-blankapp/helpers/authz"
	_ "beego-blankapp/routers"
	"github.com/astaxie/beego"
	"github.com/casbin/casbin"
	"iptelcrm/helpers/authSession"
)

func main() {
	authSession.Init()

	e := casbin.NewEnforcer("conf/authz_model.conf", "conf/authz_policy.csv")
	e.EnableLog(true)
	beego.InsertFilter("*", beego.BeforeRouter, authz.NewAuthorizer(e))
	beego.Run()
}

