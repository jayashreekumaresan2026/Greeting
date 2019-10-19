package main

import (
	"fmt"
	"bufio"
	"os"
)

var name string;

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name = scanner.Text()
		break 
	}
	return name 
}

func main(){
	fmt.Println("Hello! What is your name?")
	getName()
	fmt.Println("Hi! " + name + ". It's nice to meet you!")
}
