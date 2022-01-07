/*
	Author : Colton Smith
	Email  : colton.smith.ai@gmail.com
	Github : https://github.com/colton-smith-ai
	Date   : January 2022
*/

package watermark

import "fmt"

func Sign(date string) (sig string) {
	sig = "\nWatermark @colton.smith.ai\n" +
		"* Author : Colton Smith\n" +
		"* Email  : colton.smith.ai@gmail.com\n" +
		"* Github : https://github.com/colton-smith-ai\n" +
		"* Date   : " + date + "\n\n"
	fmt.Println(sig)
	return
}
