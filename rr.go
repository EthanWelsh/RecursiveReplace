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

	
	files := getAllFilesUnderDir(dirName)

	for i:= range files {
		replace(files[i], old, new)
	}
	
}

func replace(filepath string, old string, new string) {

	arr, _ := ioutil.ReadFile(filepath)

	s := string(arr)
	s = strings.Replace(s, old, new, -1)

	ioutil.WriteFile(filepath, []byte(s), 0744)
}

func getAllFilesUnderDir(dir string) []string {

	s := exploreDir(dir)

	return strings.Split(s, " ")
}

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
