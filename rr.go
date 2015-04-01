package main

import (
	"io/ioutil"
	"strings"
	"os"
)

func main() {

	args := os.Args[1:]

	old := args[0]
	new := args[1]
	dirName := args[2]

	// get a list of all files underneath the given one
	files := getAllFilesUnderDir(dirName)

	// for every file in this list...
	for i:= range files {
		replace(files[i], old, new) // replace all instances of the string old with new.
	}
}

// Perform the replace operation. Replace all instances of string 'old' with 'new' in and below the given 'filepath'
func replace(filepath string, old string, new string) {

	arr, _ := ioutil.ReadFile(filepath)

	s := string(arr)
	s = strings.Replace(s, old, new, -1)

	ioutil.WriteFile(filepath, []byte(s), 0744)
}

// Given a directory, will return an array of filepaths to all files and directories under the given one.
func getAllFilesUnderDir(dir string) []string {

	s := exploreDir(dir)

	return strings.Split(s, " ")
}

// Given a directory, will return a space seperated list of every file in every level underneath that directory
func exploreDir(dir string) string {
	
	file := ""
	a, _ := ioutil.ReadDir(dir)
	
	for i:= range a {
		if file != "" {
			file += " "
		}
		
		if a[i].Mode().IsDir() {
			file += exploreDir(dir + "/" + a[i].Name())
		} else {
			file += dir + "/" + a[i].Name()
		}
	}

	return file	
}
