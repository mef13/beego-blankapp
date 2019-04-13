package controllers

import (
	"github.com/astaxie/beego"
	"iptelcrm/helpers/authSession"
	"iptelcrm/models/auth"
)

type LoginController struct {
	beego.Controller
}

type LogoutController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["title"] = "Login"
	c.TplName = "login.html"

}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" && password == "" {
		c.Get()
		return
	}

	if auth.Auth(username, password) {
		sess, err := authSession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		if err != nil {
			panic(err)
		}
		defer sess.SessionRelease(c.Ctx.ResponseWriter)
		err = sess.Set("username", username)
		err = sess.Set("IsAutorized", true)
		if err != nil {
			panic(err)
		}
		c.Data["isAutorized"] = sess.Get("IsAutorized")
		c.CruSession = sess
		c.Redirect("/", 301)
	}


	c.Data["title"] = "Login"
	c.TplName = "index.tpl"

	c.Get()

}

func (c *LogoutController) Get(){
	sess, err := authSession.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		panic(err)
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	err = sess.Flush()
	if err != nil {
		panic(err)
	}
	c.Redirect("/login",301)
}


