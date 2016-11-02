package main

import "fmt"

func main() {
	for i := 100000000; i < 200000000; i++ {
				fmt.Println("This is a stdout, num:", i)
	}
}

