package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// There is one base Url
// if it doesn't have an image, then, you're out of luck until it gets added
var baseURLs = []string{"https://art.gametdb.com/ds/coverDS"}

// downloadCover downloads the cover
func downloadCover(gameID string) bool {
	regionLetter := gameID[3]
	if string(regionLetter) == "A" ||
		string(regionLetter) == "B" || string(regionLetter) == "G" {
		fmt.Println("Heya, please post an issue on the github repo " +
			"and make sure to provide the following ID: " + gameID)
		return false
	}
	if r, ok := region[string(regionLetter)]; ok {
		if !downloadBMP(gameID, r) {
			fmt.Println("Oh sorry, it looks like I'm missing the boxart for: " + gameID + " please post an issue on the github repo with the game ID and I'll see if I can get the boxart")
			// write to the missing games text file
			// if the log doesn't exist though, create it
			l, err := os.OpenFile("missing.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Can't create missing.txt file!")
				return false
			}
			if _, err := l.Write([]byte("Missing gameID: " + gameID + "  \n")); err != nil {
				fmt.Println("Can't write to missing.txt file!")
				return false
			}

			if err := l.Close(); err != nil {
				fmt.Println("Can't close missing.txt file!")
				return false
			}
			return false
		}
		return true
	}
	fmt.Println("Looks like this wasn't a ROM and slipped past my checks, whoops, GAME ID: " + gameID)
	return false
}

// downloadPNG will download the cover/box art
func downloadBMP(gameID, rg string) bool {
	fmt.Println("Attempting to download BMP cover/box art for " + gameID + ", please wait!")
	for _, url := range baseURLs {
		// first we send a HEAD request to see if the file exists
		resp, err := http.Head(url + "/" + rg + "/" + gameID + ".bmp")
		// if there's an error, go to the next loop
		if err != nil {
			continue
		}
		// if we don't got the OK, then continue to the next loop
		if resp.StatusCode != http.StatusOK {
			continue
		}
		// if we make it to this point, then there's a file there and we should
		// download it!
		resp, err = http.Get(url + "/" + rg + "/" + gameID + ".bmp")
		// if there's an error, go to the next loop
		if err != nil {
			continue
		}
		boxart, err := os.Create("boxart/" + gameID + ".bmp")
		if err != nil {
			fmt.Println("Looks like we couldn't create a file, I'm gonna exit the program here just to avoid any bigger problems")
			fmt.Println(err.Error())
			os.Exit(3)
		}
		_, err = io.Copy(boxart, resp.Body)
		if err != nil {
			fmt.Println("Looks like we couldn't copy the image contents to a file, I'm gonna exit the program here just to avoid any bigger problems")
			os.Exit(3)
		}
		// close the file
		boxart.Close()
		// return true because if we get to this part of the loop, then we've got the file
		return true
	}
	return false
}
