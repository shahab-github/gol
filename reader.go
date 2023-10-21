package main

import (
	"bufio"
	"fmt"
	"os"
)

func Test() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full name: ")
	fullName, _ := reader.ReadString('\n')

	fmt.Printf("Hello %v", fullName)
}
