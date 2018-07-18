package main

import (
	"fmt"
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
		pngExists := downloadPNG(r)
		// if the png failed to download
		if !pngExists {
			// then attempt to download a jpeg of the covert art
			jpegExists := downloadJPEG(r)
			// if the jpeg failed to download
			if !jpegExists {
				fmt.Println("Oh dear, it looks like we couldn't get the box/cover art for " + gameID + " sorry about this!")
			} else {
				convertToBMP(gameID, "jpeg")
			}
		} else {
			convertToBMP(gameID, "png")
		}
		return
	}
}

// downloadPNG will download a JPEG version of the cover art
func downloadJPEG(rg string) bool {
	fmt.Println("Attempting to download JPEG covert art, please wait!")
	for _, url := range baseURLs {
		fmt.Println(url)
	}
	return false
}

// downloadPNG will download a PNG version of the cover art
func downloadPNG(rg string) bool {
	fmt.Println("Attempting to download PNG covert art, please wait!")
	for _, url := range baseURLs {
		fmt.Println(url)
	}
	return false
}
