package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		return
	}
	defer file.Close()
	//Get the file size
	stat, err := file.Stat()
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError

	if err != nil {
		return
	}
	fmt.Println(stat.Size())
	//Read the file
	sl := make([]byte, stat.Size())
	_, err = file.Read(sl)
	// Read reads up to len(b) bytes from the File.
	// It returns the number of bytes read and any error encountered.
	// At end of file, Read returns 0, io.EOF.

	if err != nil {
		return
	}
	str := string(sl)
	fmt.Println(str)

	// Note there is a shorter way to do this as follows:
	// USING IOUTIL
	fileContents, err := ioutil.ReadFile("text.txt")
	// ReadFile reads the file named by filename and returns the contents.
	// A successful call returns err == nil, not err == EOF.
	// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.
	if err != nil {
		return
	}
	str2 := string(fileContents)
	fmt.Println()
	fmt.Printf("Same file as above reads as follows:\n%s\n", str2)
}
