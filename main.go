package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	var age int
	scanner.Scan()
	ageStr := scanner.Text()
	age, _ = strconv.Atoi(ageStr)
	fmt.Printf("Hi, %s! You are %d years old.\n", name, age)

}
