package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// There are 4 base urls
// if one doesn't have an image, the next one is used
var baseURLs = []string{"https://art.gametdb.com/ds/cover/",
	"https://art.gametdb.com/ds/coverS/",
	"https://art.gametdb.com/ds/coverM/",
	"https://art.gametdb.com/ds/coverHD/"}

// downloadCover downloads the cover
func downloadCover(gameID string) {
	regionLetter := gameID[3]
	if string(regionLetter) == "A" ||
		string(regionLetter) == "B" || string(regionLetter) == "G" {
		fmt.Println("Heya, please post an issue on the github repo " +
			"and make sure to provide the following ID: " + gameID)
		return
	}
	if r, ok := region[string(regionLetter)]; ok {
		// first attempt to download a PNG of the cover art
		pngExists := downloadPNG(r, gameID)
		// if the png failed to download
		if !pngExists {
			// then attempt to download a jpeg of the covert art
			jpegExists := downloadJPEG(r, gameID)
			// if the jpeg failed to download
			if !jpegExists {
				fmt.Println("Oh dear, it looks like we couldn't get the box/cover art for " + gameID + " sorry about this!")
			} else {
				fmt.Println("Box art found! Converting to bmp now!")
				convertToBMP(gameID, "jpeg")
			}
		} else {
			fmt.Println("Box art found! Converting to bmp now!")
			convertToBMP(gameID, "png")
		}
		return
	}
	fmt.Println("Looks like this wasn't a ROM and slipped past my checks, whoops")
	return
}

// downloadPNG will download a JPEG version of the cover art
func downloadJPEG(rg, gameID string) bool {
	fmt.Println("Attempting to download JPEG cover art for " + gameID + ", please wait!")
	for _, url := range baseURLs {
		// first we send a HEAD request to see if the file exists
		resp, err := http.Head(url + "/" + rg + "/" + gameID + ".jpeg")
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
		resp, err = http.Get(url + "/" + rg + "/" + gameID + ".jpeg")
		// if there's an error, go to the next loop
		if err != nil {
			continue
		}
		tmpFile, err := os.Create("tmp.jpeg")
		if err != nil {
			fmt.Println("Looks like we couldn't create a file, I'm gonna exit the program here just to avoid any bigger problems")
			os.Exit(3)
		}
		_, err = io.Copy(tmpFile, resp.Body)
		if err != nil {
			fmt.Println("Looks like we couldn't copy the image contents to a file, I'm gonna exit the program here just to avoid any bigger problems")
			os.Exit(3)
		}
		// close the file
		tmpFile.Close()
		// return true because if we get to this part of the loop, then we've got the file
		return true
	}
	return false
}

// downloadPNG will download a PNG version of the cover art
func downloadPNG(rg, gameID string) bool {
	fmt.Println("Attempting to download PNG cover art for " + gameID + ", please wait!")
	for _, url := range baseURLs {
		// first we send a HEAD request to see if the file exists
		resp, err := http.Head(url + "/" + rg + "/" + gameID + ".png")
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
		resp, err = http.Get(url + "/" + rg + "/" + gameID + ".png")
		// if there's an error, go to the next loop
		if err != nil {
			continue
		}
		tmpFile, err := os.Create("tmp.png")
		if err != nil {
			fmt.Println("Looks like we couldn't create a file, I'm gonna exit the program here just to avoid any bigger problems")
			os.Exit(3)
		}
		_, err = io.Copy(tmpFile, resp.Body)
		if err != nil {
			fmt.Println("Looks like we couldn't copy the image contents to a file, I'm gonna exit the program here just to avoid any bigger problems")
			os.Exit(3)
		}
		// close the file
		tmpFile.Close()
		// return true because if we get to this part of the loop, then we've got the file
		return true
	}
	return false
}
