/*
	Author : Colton Smith
	Title  : Data Science Engineer
	Email  : colton.smith.ai@gmail.com
	Github : https://github.com/colton-smith-ai
	Date   : January 2022
*/

package signature

import "fmt"

func Sign(date string) (sig string) {
	sig = "\nSignature @colton.smith.ai\n" +
		"\tAuthor : Colton Smith\n" +
		"\tTitle  : Data Science Engineer\n" +
		"\tEmail  : colton.smith.ai@gmail.com\n" +
		"\tGithub : https://github.com/colton-smith-ai\n" +
		"\tDate   : " + date + "\n\n"
	fmt.Println(sig)
	return
}
