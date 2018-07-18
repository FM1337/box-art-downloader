package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

// This offset array gets us our game id
var offsets = []uint8{0xC, 0x10}

func main() {
	// check to make sure we've been provided a file or a folder
	if len(os.Args) > 1 {
		// make the maps before reading the game
		makeMap()
		// read the game
		id := readOffset(os.Args[1])
		// check to make sure id isn't empty
		if !validate(id) {
			// exit with 2 if it is
			fmt.Println("This is not a valid .nds rom!")
			os.Exit(2)
		}
		downloadCover(id)
		os.Exit(0)
	}
	fmt.Println("You must provide a rom or a folder of roms!")
	os.Exit(1)
}

func readOffset(file string) string {
	game, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	// return the game
	return string(game[offsets[0]:offsets[1]])
}

// validate will validate the game ID to make sure it's valid
func validate(gameID string) bool {
	x := regexp.MustCompile(`[A-Z][A-Z][A-Z][A-Z]`)
	return x.MatchString(gameID)
}
