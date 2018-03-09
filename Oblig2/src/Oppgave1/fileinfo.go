package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)



func main() {
	ReadFile()
	Directory()
	Regular()
	Unix()
	Append()
	Devicefile()
}

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
	fmt.Printf("\nLength: %e gigabytes\n", f/1e9)
}


func Directory() {
	fi, err := os.Lstat("fileinfo.go")
	if err != nil {
		log.Fatal(err)}

	switch mode11 := fi.Mode; {
	case mode11().IsDir() == true:
		fmt.Printf("\nThe directory does not exist \n")
	case mode11().IsDir() == false:
		fmt.Printf("\nThe directory exist")
	}
}

func Regular() {
	fi, err := os.Lstat("fileinfo.go")
	if err != nil {
		log.Fatal(err)}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Printf("\nIs a regular file")
	case !mode.IsRegular():
		fmt.Printf("\nNot a regular file")
	}
}

func Unix () {
	fi, err := os.Lstat("fileinfo.go")
	if err != nil {
		log.Fatal(err)}

	switch mode1 := fi.Mode(); {
	case mode1&os.ModePerm == 0777:
		fmt.Printf("\nHas Unix permission bits")
	case mode1&os.ModePerm != 0770:
		fmt.Printf("\nHas NOT Unix permission bits")

	}
}

func Append() {
	fi, err := os.Lstat("fileinfo.go")
	if err != nil {
		log.Fatal(err)}

	switch mode2 := fi.Mode(); {
	case mode2&os.ModeAppend == 0:
		fmt.Printf("\nIs append only")
	case mode2&os.ModeAppend != 0:
		fmt.Printf("\nIs NOT append only")
	}

}

func Devicefile()  {
	fi, err := os.Lstat("fileinfo.go")
	if err != nil {
		log.Fatal(err)}

	switch mode3 := fi.Mode(); {
	case mode3&os.ModeDevice == 0:
		fmt.Printf("\nIs not a device")
	case mode3&os.ModeDevice != 0:
		fmt.Printf("\nIs a device")
	}
}