package main

import (
	"errors"
	"fmt"
)

func main() {
	defer catchErr()
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword2(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

// Function Recover
func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly123")
	}
}

func validPassword2(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password has to have more than 4 characters")
	}
	return "Valid Password", nil
}
