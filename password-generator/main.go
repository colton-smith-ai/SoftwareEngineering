/*
	Author : Colton Smith
	Title  : Data Science Engineer
	Date   : December 2021
	Email  : colton.smith.ai@gmail.com
	Github : https://github.com/colton-smith-ai
*/

package main

import (
	"colton-smith-ai/SoftwareEngineering/password-generator/newpass"
	"fmt"
)

func main() {
	randomPassword := newpass.NewPassword()
	fmt.Println(randomPassword)
}
