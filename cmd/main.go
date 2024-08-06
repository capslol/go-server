package main

import "fmt"

func main() {

	counter := 0

	for {
		counter++
		fmt.Println(counter)
		if counter == 100 {
			break
		}

	}
}
