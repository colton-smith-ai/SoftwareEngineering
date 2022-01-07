package simpleio

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func UserInput(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	PrinterLine(msg)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func HandleErrors(err error) {
	fmt.Print("An error has occured! : ")
	fmt.Println(err)
	log.Fatalln(err)
}

func Printer(message string) {
	fmt.Println(message)
}

func PrinterLine(message string) {
	fmt.Print(message)
}
