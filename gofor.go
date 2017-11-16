package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Weekday() == time.Thursday)

}

func test() {

}

func what(i interface{}) {

	switch t := i.(type) {
	case bool:
		fmt.Println("You are boolean type")
		fmt.Println("You are ", t)
	case int:
		fmt.Println("You are integer")
	case string:
		fmt.Println("you are string")
	default:
		fmt.Println("I dont really know the tyoe")
		fmt.Println("You are ", t)
	}

}

//let me write a function that can take any type of argument
