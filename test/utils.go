package test

import (
	"log"
	"os"
)

func CheckDirs(dirNames []string) {
	for _, dirName := range dirNames {
		ensureDir(dirName)
	}
}

func ensureDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %s", err)
		}
	} else if err != nil {
		log.Fatalf("Error checking if directory exists: %s", err)
	}
}
