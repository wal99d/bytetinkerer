package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wal99d/listComprehension/bytetinkerer"
)

func main() {
	file, err := os.Open("dummyFile")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer file.Close()

	filteredBytes, err := bytetinkerer.From(
		file,
		bytetinkerer.Remove([]byte("GefAliases()")),     //remove some noise from our data
		bytetinkerer.Extract([]byte("GefTmuxSetup()")),  //extract new list from our old data
		bytetinkerer.Stamp(time.Now().Local().String()), //Put a timestamp on our data
	)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(filteredBytes.ConvertToString()) //Show us the filtered Bytes :)
}
