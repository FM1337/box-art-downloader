package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// convertToBMP converts the image to bmp and saves it
func convertToBMP(gameID, fileType string) {
	args := []string{"tmp." + fileType, "-resize", "128x115", "-define", "bmp:subtype=RGB555", gameID + ".bmp"}
	err := exec.Command("./linux/magick", args...).Run()

	if err != nil {
		log.Fatal(err)
	}

	// once the conversion is done, we should delete the tmp file
	err = os.Remove("tmp." + fileType)
	if err != nil {
		fmt.Println("There was an error deleting the tmp file? Did you already delete it?")
	}
	fmt.Println("Conversion of box art is complete!")
}
