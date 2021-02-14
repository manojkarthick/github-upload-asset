package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Check whether the given file exists on the filesystem
func CheckFileExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Errorf("Path %s does not exist", path)
		return err
	}
	return nil
}

// Get the base name from the path
func GetBaseName(path string) string {
	return filepath.Base(path)
}
