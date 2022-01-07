/*
	Author : Colton Smith
	Title  : Data Science Engineer
	Email  : colton.smith.ai@gmail.com
	Github : https://github.com/colton-smith-ai
	Date   : January 2022
*/

package author_watermark

import "fmt"

func Sign(date string) (sig string) {
	sig = "\nSignature @colton.smith.ai\n" +
		"* Author : Colton Smith\n" +
		"* Title  : Data Science Engineer\n" +
		"* Email  : colton.smith.ai@gmail.com\n" +
		"* Github : https://github.com/colton-smith-ai\n" +
		"* Date   : " + date + "\n\n"
	fmt.Println(sig)
	return
}
