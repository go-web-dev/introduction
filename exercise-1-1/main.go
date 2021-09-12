package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf(
		"My name is %s, I'm %d years old\n",
		"Steve Hook",
		time.Now().Year()-1995,
	)
}
