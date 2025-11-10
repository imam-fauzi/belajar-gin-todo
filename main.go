package main

import "fmt"

func main() {
	for r := 1; r <= 5; r++ {
		for c := 1; c <= r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
