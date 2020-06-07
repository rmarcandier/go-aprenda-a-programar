package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	senha := "20julho1980"

	bs, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)

}
