package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

// This offset array gets us our game id
var offsets = []uint8{0xC, 0x10}

func main() {
	// check to make sure we've been provided a file or a folder
	if len(os.Args) > 1 {
		// let's check to see if the boxart directory exists and create it if it doesn't
		if _, err := os.Stat("./boxart"); os.IsNotExist(err) {
			// if it doesn't, then create the directory!
			os.Mkdir("boxart", 0700)
		}
		// check to see if we're reading a directory or a file
		file, err := os.Stat(os.Args[1])
		if err != nil {
			fmt.Println("There's a problem reading your file/folder!")
			os.Exit(4)
		}
		// if we're reading a directory
		if file.IsDir() {
			games, err := ioutil.ReadDir(os.Args[1])
			if err != nil {
				fmt.Println("There's a problem reading your files in your folder!")
				os.Exit(5)
			}
			// when wait hits 10, we'll want to wait a few seconds before
			// continuing so that we don't get blocked
			wait := 0
			successful := 0
			// looping time!
			for _, game := range games {
				if strings.HasSuffix(game.Name(), "nds") {
					// get the id
					id := readOffset(os.Args[1] + "/" + game.Name())
					// make sure the id isn't empty
					if id == "" {
						//fmt.Println(game.Name() + " is not a valid nds rom!")
						// instead of exiting, just continue to the next loop
						continue
					}
					// if we do have an ID, next we check the validation
					if !validate(id) {
						// exit with 2 if it is
						//fmt.Println(game.Name() + " is not a valid nds rom!")
						// instead of exiting, just continue to the next loop
						continue
					}
					// if we make it to this part of the loop, we've passed the main checks!
					// so now we want to actually begin downloading the cover!
					success := downloadCover(id)
					if success {
						successful++
					}
					wait++
					if wait == 10 {
						fmt.Println("Waiting 5 seconds before continuing so we don't get blocked by the server!")
						time.Sleep(5 * time.Second)
						wait = 0
					}
				} else {
					//fmt.Println(game.Name() + " is not a valid nds rom!")
				}
			}
			// once the loop is finished
			// exit the script.
			fmt.Println(fmt.Sprintf("We were able to download %d/%d of the boxart for your games!", successful, len(games)))
			if successful == len(games) {
				fmt.Println("Awesome! We were able to get boxart for all your games!")
			}
			fmt.Println("Thank you for using the box art downloader by Allen (FM1337)!")
			fmt.Println("Have a nice day!")
			os.Exit(0)
		} else {
			// read the game
			id := readOffset(os.Args[1])
			// check to make sure id isn't empty
			if id == "" {
				fmt.Println("This is not a valid nds rom!")
				os.Exit(2)
			}
			if !validate(id) {
				// exit with 2 if it is
				fmt.Println("This is not a valid nds rom!")
				os.Exit(2)
			}
			success := downloadCover(id)
			if !success {
				fmt.Println("I'm sorry that we weren't able to find your boxart, don't forget to post an issue on the github repo and I'll do my best to get the art added!")
				os.Exit(0)
			}
			fmt.Println("Box art has been downloaded!")
			fmt.Println("Thank you for using the box art downloader by Allen (FM1337)!")
			fmt.Println("Have a nice day!")
			os.Exit(0)
		}
	}
	fmt.Println("You must provide a rom or a folder of roms!")
	os.Exit(1)
}

func readOffset(file string) string {
	game, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	// return the game
	return string(game[offsets[0]:offsets[1]])
}

// validate will validate the game ID to make sure it's valid
func validate(gameID string) bool {
	x := regexp.MustCompile(`[A-Z0-9][A-Z0-9][A-Z0-9][A-Z40-9]`)
	return x.MatchString(gameID)
}
