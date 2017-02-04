package B

import . "../D"

type B struct {

}

func (b* B)B1(){
	println("B1 func")
}

func (a *B)B2(){
	println("B2 func ")
}

func UseA(a AI){
	a.A1()
}

func GetB()*B{
	b := new(B)
	return b
}
