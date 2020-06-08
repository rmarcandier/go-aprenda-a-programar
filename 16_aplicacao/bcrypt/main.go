package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	senha := "09Agosto1979"

	bs, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	fmt.Println(bcrypt.CompareHashAndPassword(bs, []byte(senha)))

}
