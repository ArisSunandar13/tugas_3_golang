package main

import "fmt"

func main() {
	var buah = []string{"apel", "jeruk", "anggur", "melon"}

	buah = append(buah, "pepaya")

	for i := 0; i < len(buah); i++ {
		fmt.Printf("Element ke-%d : %s", i, buah[i])
		fmt.Println()
	}
}
