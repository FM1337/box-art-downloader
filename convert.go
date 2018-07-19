package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// convertToBMP converts the image to bmp and saves it
func convertToBMP(gameID, fileType string) {
	args := []string{"tmp." + fileType, "-resize", "128x115", "-define", "bmp:subtype=RGB555", gameID + ".bmp"}
	// Check what os we're running
	OS := runtime.GOOS
	switch OS {
	case "windows":
		fmt.Println("Windows isn't supported yet, sorry!")
		os.Exit(6)
	case "linux":
		err := exec.Command("./linux/magick", args...).Run()

		if err != nil {
			log.Fatal(err)
		}
	case "darwin":
		fmt.Println("OS X isn't supported yet, sorry!")
		os.Exit(6)
	}

	// once the conversion is done, we should delete the tmp file
	err := os.Remove("tmp." + fileType)
	if err != nil {
		fmt.Println("There was an error deleting the tmp file? Did you already delete it?")
	}
	fmt.Println(gameID + "'s box art has been converted to bmp!")
}
