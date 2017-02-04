package A

import . "../B"

type A struct {

}

func (a* A)A1(){
	println("A1 func")
}

func (a *A)A2(){
	println("A2 func ")
}

func UseB(){
	b := GetB()
	b.B1()
}

func GetA()*A{
	a := new(A)
	return a
}

/*
	A --> B
	B --> D
	C --> AB




 */



