package authSession

import (
	"github.com/astaxie/beego/session"

)

var (
	GlobalSessions *session.Manager
)

func Init() {
	conf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
		EnableSetCookie: true,
	}

	GlobalSessions, _ = session.NewManager("memory", conf)
	go GlobalSessions.GC()
}
