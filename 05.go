package main

import (
	"fmt" 
	"time"
)

var a = 1;

func demo(){
	var a = 3;
	fmt.Println(a);
}


func main(){

	demo()
	fmt.Println(a)

	if a == 1{
		fmt.Println("Hello")
	}else if a == 2{
		fmt.Println("Hi")
	}else {
		fmt.Println("Hi 1")
	}

	today := time.Now().Weekday()
	fmt.Println(time.Now())
	fmt.Println(time.Thursday)
	fmt.Println(today)
	switch time.Thursday{
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Friday")
	}

}