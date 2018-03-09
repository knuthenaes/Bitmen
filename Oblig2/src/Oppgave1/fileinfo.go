package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)


func ReadFile() {
	data, err := ioutil.ReadFile("fileinfo.go")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len(data))
	i := len(data)
	f := float64(i)

	fmt.Printf("\nLength: %f kilobytes", f/1000)
	fmt.Printf("\nLength: %f megabytes", f/1e6)
	fmt.Printf("\nLength: %e gigabytes", f/1e9)
}

func Directory() {
	file, _ := os.Stat("fileinfo.go")
	file.IsDir()
	if _, err := os.Stat("fileinfo.go"); os.IsNotExist(err){
		fmt.Printf("\nThe directory does not exist \n")
	}
	if _, err := os.Stat("fileinfo.go"); err == nil {
		fmt.Printf("\n")
		fmt.Printf("\nThe directory exist")
	}
}

func regular() {

}