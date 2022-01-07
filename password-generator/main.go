/*
	Author : Colton Smith
	Title  : Data Science Engineer
	Email  : colton.smith.ai@gmail.com
	Github : https://github.com/colton-smith-ai
	Date   : December 2021
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
