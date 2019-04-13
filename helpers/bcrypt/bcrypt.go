package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)


func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		//fmt.Println(arg)
		password := []byte(arg)
		ncryptedPassword, err := bcrypt.GenerateFromPassword(password, 13)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(ncryptedPassword))
	}
}
