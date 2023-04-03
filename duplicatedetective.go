package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Check that the user provided a directory to search
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <directory>\n", os.Args[0])
		os.Exit(1)
	}
	// Get the directory to search
	directory := os.Args[1]

	// Create a map to hold file hashes and paths
	fileMap := make(map[string]string)

	// Walk through the directory and its subdirectories
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		// If there was an error, print it and skip this file/directory
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// If this isn't a regular file, skip it
		if !info.Mode().IsRegular() {
			return nil
		}

		// Read the file contents
		fileContents, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// Compute the MD5 hash of the file contents
		hash := md5.Sum(fileContents)
		hashString := fmt.Sprintf("%x", hash)

		// If we've seen this hash before, print a message saying that the file is a duplicate
		if existingPath, ok := fileMap[hashString]; ok {
			fmt.Printf("%s is a duplicate of %s\n", path, existingPath)
		} else {
			// Otherwise, add the hash and path to the map
			fileMap[hashString] = path
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
