package main

import "fmt"

func add(a int, b int) int{
	return a +   b
}

func add1(a, b int) int{
	return a + b
}

func twoFunction(a, b int) (int, int){
	return (a + b), (a * b)
}


func twoFunction1(a, b int) (check, check2 int){
	check = (a + b)
	check2 = (a * b)
	return 
}


func main4(){

	a := 1
	b := 2

	fmt.Println(add(a,b))
	fmt.Println(add1(a,b))
	fmt.Println(twoFunction(a,b))
	fmt.Println(twoFunction1(a,b))
}