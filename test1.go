package main

import "fmt"

func f()(err string) {

	return err
}

func main() {
	e := f()
	if e != nil {
		fmt.Println("e is not nil")
	}else {
		fmt.Println("e is nil")
	}

}
