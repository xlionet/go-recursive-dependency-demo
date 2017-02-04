package main

func Add1()string{

	str1 := "123456789"
	str2 := "987654321"

	str := str1 + str2

	return str
}

func Add2()string{
	str1 := "123456789"
	str2 := "987654321"
	copy(str)
}
