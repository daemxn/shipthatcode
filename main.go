package main

import (
	"fmt"
	"strconv"
)

func main() {
	var line string
	fmt.Scan(&line)

	n, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("bad")
	} else {
		fmt.Printf("ok %d\n", n)
	}

}
