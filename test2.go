package main

import "fmt"

func f1(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")


}

func main() {

	f1()
}