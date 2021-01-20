package main

// go run flag.go | xclip

import (
	"fmt"
	"io/ioutil"
)

func main() {
	cat, _ := ioutil.ReadFile("cattos.jpg")
	kit, _ := ioutil.ReadFile("kitters.jpg")
	for i := 0; i < len(cat); i++ {
		k := kit[i]
		c := cat[i]
		if k != c {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
