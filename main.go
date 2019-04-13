package main

import (
	"beego-blankapp/helpers/authSession"
	"beego-blankapp/helpers/authz"
	_ "beego-blankapp/routers"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/casbin/casbin"
	"github.com/getsentry/raven-go"
)

func init()  {
	sentryDSN := beego.AppConfig.String("sentryDSN")
	if sentryDSN != "" {
		raven.SetDSN(sentryDSN)
		if beego.BConfig.RecoverPanic {
			initRecover()
		}
	}
}

func initRecover() {
	originRecover := beego.BConfig.RecoverFunc
	beego.BConfig.RecoverFunc = func(ctx *context.Context) {
		defer originRecover(ctx)
		if err := recover(); err != nil {
			errStr := fmt.Sprint(err)
			packet := raven.NewPacket(errStr, raven.NewException(errors.New(errStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(ctx.Request))
			raven.Capture(packet, nil)
			panic(err)
		}
	}
}


func main() {
	authSession.Init()

	e := casbin.NewEnforcer("conf/authz_model.conf", "conf/authz_policy.csv")
	e.EnableLog(true)
	beego.InsertFilter("*", beego.BeforeRouter, authz.NewAuthorizer(e))
	beego.Run()
}

