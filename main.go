package main

import "fmt"

func tambah(a, b int) int { return a + b }
func kurang(a, b int) int { return a - b }

func main() {
	fmt.Println(tambah(2, 5))
	fmt.Println(kurang(8, 3))
}
