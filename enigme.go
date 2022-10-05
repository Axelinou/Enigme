package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("C:/Users/shnou/OneDrive/Documents/Ytrac/file/Hangman.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ReadFile(string(file))[0])
	fmt.Println(len(ReadFile(string(file))) - 1)
	fmt.Println()
}

func ReadFile(file string) []string {
	var arr []string
	var tmp string

	for _, _byte := range file {

		if _byte == '\n' {
			arr = append(arr, tmp)
			tmp = ""
		} else {
			tmp += string(_byte)
		}
	}

	if tmp == "before" {
		fmt.Println(string(file)[14] + 1)
	}
	return arr
}
