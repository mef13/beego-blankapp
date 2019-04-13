# beego-blankapp 
beego-blankapp is a beego based template with some libs integrations.


### Tech
 - Authorization based on beego sessions
 - Passwords hashed by bcrypt
 - ReactJS
 - Websocket
 - Casbin
 - Sentry for Go
 
 
### Installation
    go get github.com/astaxie/beego
    go get github.com/casbin/casbin
    go get golang.org/x/crypto/bcrypt
    go get github.com/gorilla/websocket
    go get github.com/getsentry/raven-go


### Authorization
Users and hashed password in conf/users.conf.

Get hash from new password: 
```sh
$ go run helpers/bcrypt/bcrypt.go newpassword
```


### Todos
 - Add Sentry for React
 - Save users in PostgreSQL

