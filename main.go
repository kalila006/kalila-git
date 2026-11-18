package main

import "fmt"

func main() {

	Hari := 1
	nominal := 1
	var total int
	for Hari <= 7 {
		fmt.Println("hari ke\n ", Hari)
		Hari += 1

		fmt.Print("nominal ")
		fmt.Scan(&nominal)
		total += nominal
	}
	fmt.Println("Kelua")
	fmt.Println("total data ", total)
}
