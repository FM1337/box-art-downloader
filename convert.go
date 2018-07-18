package main

import (
	"log"
	"os/exec"
)

// convertToBMP converts the image to bmp and saves it
func convertToBMP(gameID, fileType string) {
	args := []string{"tmp." + fileType, "-resize", "128x115", "-define", "bmp:subtype=RGB555", gameID + ".bmp"}
	err := exec.Command("./linux/magick", args...).Run()

	if err != nil {
		log.Fatal(err)
	}
}
