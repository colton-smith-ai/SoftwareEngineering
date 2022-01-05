package newpass

import (
	"math/rand"
	"strings"
	"time"
)

const PASSWORD_LENGTH = 20

func NewPassword() string {
	// proper randomize with seed
	rand.Seed(time.Now().UnixNano())

	// create encoding obj
	ac := createCodex()

	// create string placeholder
	builder := [PASSWORD_LENGTH]string{}

	// add three symbols
	insertSymbol(3, &builder, ac.sym)

	// add three numbers
	insertRandom(3, &builder, ac.num["min"], ac.num["max"])

	// add three capitals
	insertRandom(5, &builder, ac.alphaCap["min"], ac.alphaCap["max"])

	// add random char
	randCharChoice := rand.Intn(3)
	switch randCharChoice {
	case 0:
		insertSymbol(1, &builder, ac.sym)
	case 1:
		insertRandom(1, &builder, ac.num["min"], ac.num["max"])
	case 2:
		insertRandom(1, &builder, ac.alphaCap["min"], ac.alphaCap["max"])
	default:
		insertRandom(1, &builder, ac.alphaLow["min"], ac.alphaLow["max"])
	}

	// fill remaining with lower case
	for i := 0; i < len(builder); i++ {
		// check for empty/available spaces
		if builder[i] == "" {
			// get random ascii int
			randAsciiInt := randomRange(ac.alphaLow["min"], ac.alphaLow["max"])
			// convert ascii int to char/string
			asciiToChar := string(rune(randAsciiInt))
			// insert random char in password
			builder[i] = asciiToChar
		}
	}

	// convert placeholder array to string
	return strings.Join(builder[:], "")
}

func randomRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func insertRandom(amountToInsert int, arrayToUpdate *[PASSWORD_LENGTH]string, rangeMin int, rangeMax int) {
	for i := 0; i < amountToInsert; i++ {
		// get random char ascii int
		randomAsciiInt := randomRange(rangeMin, rangeMax)
		// convert ascii int into string representation
		asciiToChar := string(rune(randomAsciiInt))
		// find random index to insert char into password
		randIndex := rand.Intn(len(arrayToUpdate))
		// check index location is free in password
		for arrayToUpdate[randIndex] != "" {
			// find random index to place symbol in password
			randIndex = rand.Intn(len(arrayToUpdate))
		}
		// insert symbol into random location in password
		arrayToUpdate[randIndex] = asciiToChar
	}
}

func insertSymbol(amountToInsert int, arrayToUpdate *[PASSWORD_LENGTH]string, symbolArray []int) {
	for i := 0; i < amountToInsert; i++ {
		// random index for symbol array
		randIndex := rand.Intn(len(symbolArray))
		// from random index, get symbol ascii int value
		symbolInt := symbolArray[randIndex]
		// convert ascii int into string representation
		symbolChar := string(rune(symbolInt))
		// find random index to place symbol in password
		randIndex = rand.Intn(len(arrayToUpdate))
		// check index location is free in password
		for arrayToUpdate[randIndex] != "" {
			// find random index to place symbol in password
			randIndex = rand.Intn(len(arrayToUpdate))
		}
		// insert symbol into random location in password
		arrayToUpdate[randIndex] = symbolChar
	}
}

type asciiCodex struct {
	/*
		~~ASCII CODEX~~
		Special Symbols (!,@,#,$...) = 33,35-38,43,60,62-64
		Numbers (0-9) = 48-57
		Lower Alphabet (a-z) = 97-122
		Upper Alphabet (A-Z) = 65-90
	*/

	sym      []int          `default:"[33,35,36,37,38,43,60,62,63,64]"`
	num      map[string]int `default:"{\"min\": 48, \"max\": 57}"`
	alphaLow map[string]int `default:"{\"min\": 97, \"max\": 122}"`
	alphaCap map[string]int `default:"{\"min\": 65, \"max\": 90}"`
}

func createCodex() asciiCodex {
	return asciiCodex{
		sym:      []int{33, 35, 36, 37, 38, 43, 60, 62, 63, 64},
		num:      map[string]int{"min": 48, "max": 57},
		alphaLow: map[string]int{"min": 97, "max": 122},
		alphaCap: map[string]int{"min": 65, "max": 90},
	}
}
