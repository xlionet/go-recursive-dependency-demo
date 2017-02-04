package main


type A struct{
	id int
}

func (a *A)GetA()int{
	println("I am A.", a.id)
	return a.id
}

func (a *A)TrancelateToB()*B{

	return (*B)(a)
}

type B A

func (b *B)GetB()int{
	println("I am B.", b.id)
	return b.id
}

func (b *B)TrancelateToA()*A{

	return (*A)(b)
}

func main(){
	a := A{10}
	a.GetA()

	b := a.TrancelateToB()
	b.GetB()

	aa := b.TrancelateToA()
	aa.GetA()


}