package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// There are 4 base urls
// if one doesn't have an image, the next one is used
var baseURLs = []string{"https://boxart.fm1337.com/boxart"}

// downloadCover downloads the cover
func downloadCover(gameID string) bool {
	regionLetter := gameID[3]
	if string(regionLetter) == "A" ||
		string(regionLetter) == "B" || string(regionLetter) == "G" {
		fmt.Println("Heya, please post an issue on the github repo " +
			"and make sure to provide the following ID: " + gameID)
		return false
	}
	if !downloadBMP(gameID) {
		fmt.Println("Oh sorry, it looks like I'm missing the boxart for: " + gameID + " please post an issue on the github repo with the game ID and I'll see if I can get the boxart")
		return false
	}
	return true
}

// downloadPNG will download the cover/box art
func downloadBMP(gameID string) bool {
	fmt.Println("Attempting to download BMP cover/box art for " + gameID + ", please wait!")
	for _, url := range baseURLs {
		// first we send a HEAD request to see if the file exists
		resp, err := http.Head(url + "/" + gameID + ".bmp")
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
		resp, err = http.Get(url + "/" + gameID + ".bmp")
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
