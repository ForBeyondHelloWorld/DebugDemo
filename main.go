package main

import "fmt"

func main() {
	data := []string{"Happy", "Day", "Every", "Day"}
	for i, item := range data {
		fmt.Println(i, item)
	}
}
