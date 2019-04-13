package auth

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"golang.org/x/crypto/bcrypt"
)

func Auth(username, password string) bool {

	iniconf, err := config.NewConfig("ini", "conf/users.conf")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return false
	}
	hpass := iniconf.String(username)
	if hpass == "" {
		fmt.Printf("Error: Incorrect username")
		return false
	}

	encryptionErr := bcrypt.CompareHashAndPassword([]byte(hpass), []byte(password))
	if encryptionErr == nil {
		return true
	}
	return false
}
