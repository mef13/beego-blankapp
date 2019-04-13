# beego-blankapp 
beego-blankapp это шаблон на основе beego с применением некоторых технологий.


### Tech
 - Авторизация на основе beego сессий
 - Хеширование пароля bcrypt
 - ReactJS
 - Websocket
 - Casbin
 
 
### Installation
    go get github.com/astaxie/beego
    go get github.com/casbin/casbin
    go get golang.org/x/crypto/bcrypt
    go get github.com/gorilla/websocket


### Authorization
Пользователи и хешированные пароли хранятся в conf/users.conf.

Получение хеша нового пароля: 
```sh
$ go run helpers/bcrypt/bcrypt.go newpassword
```


### Todos
 - Add Sentry
 - Хранение пароля в PostgreSQL

