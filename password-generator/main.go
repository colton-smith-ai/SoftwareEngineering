/*
	Author : Colton Smith
	Email  : colton.smith.ai@gmail.com
	Github : https://github.com/colton-smith-ai
	Date   : January 2022
*/

package main

import (
	"fmt"
	"password-generator/newpass"
)

func main() {
	randomPassword := newpass.NewPassword()
	fmt.Println(randomPassword)
}
