package main

import "fmt"

func main() {
	var num1 int = 5
	var num2 int = 5

	sum := num1 + num2

	fmt.Println("Resultt :", sum)


	if sum <= 10{
		fmt.Println("Result is Equal to or less then 10")
	}else{
		fmt.Println("Result is greater then 10")
	}


	for i:= 0; i<=10; i++{
		fmt.Println("value : ",i)
	}


	result := addTwoDigit(num1, num2)
	fmt.Println("Result : ",result);
}



func addTwoDigit(num1 int, num2 int) int {
	sum := num1 + num2;
	return sum
}