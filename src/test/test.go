package main

import "fmt"

func main() {
	ar := []int{1, 2, 3, 4}
	fmt.Println("%d:%d", len(ar), cap(ar))
	ar = append(ar, 3, 4, 5)
	fmt.Println("%d:%d", len(ar), cap(ar))
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Println("%d:%d", timeZone["EST"])
}
