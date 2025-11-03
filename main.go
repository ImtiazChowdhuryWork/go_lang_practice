package main

import (
	"fmt"
	"maps"
	"slices"
	"time"
)

func main() {
	// fmt.Println("This is first go program")

	// sum := add(20,10)

	// fmt.Println("Result : ",sum)


	// slicePractice();
	// mapPractice()
	// rangePractice()
	// printValueOfTheFunctionPractice(10,20)
	// myFunc := returnFunctionFromAFunction(10)

	// mes, num := myFunc("Chowdhury")
	// fmt.Println("My Message : ",mes)
	// fmt.Println("Number : ",num)


	// res := variadicFucntionPractice(2,2,2,5,20,2)
	// fmt.Println("Resulttt : ",res)


	// myOrderrr := structPractice()
	// fmt.Println("My  Order : ",myOrderrr)






// var a int = 40
// var b int = 50

// var x int = 10
// var y int = 20

// 	newVariable1 := pointerPractice(&a,&b)

// 	  newVariable2 := newVariable1(&x, &y)
// 	  fmt.Println(newVariable2);


// 	  fmt.Println("Value of A : ",a)
// 	  fmt.Println("Value of B : ",b)


//Practice : DateTime
	dateTimePractice()

}


func add (num1 int, num2 int)int{
	res := num1+num2

	return res 
}


func slicePractice (){
	var num1[]int
	fmt.Println("Num 1 : ",num1)
	fmt.Println("Is Num 1 Nil? : ", num1 == nil)

	fmt.Println("Next Make slice with \"make'\" : ")
	var num2 = make([]int,0,5)

	fmt.Println("Num 2 : ",num2)
	fmt.Println("Num 2 Length : ",len(num2))
	fmt.Println("Num 2 Capacity : ",cap(num2))

	fmt.Println("Let's use append on num2 : ")
	num2 = append(num2, 5,4,7,8,9,5,2,2,3,4,44)
	fmt.Println("Print the append value on num2 slice : ",num2)
	fmt.Println("Num 2 Length : ",len(num2))
	fmt.Println("Num 2 Capacity : ",cap(num2))



	
	var num3 = make([]int, len(num2))
	fmt.Println("Declared the num 3 slice with the length of num2 slice", num3)
	fmt.Println("Num 3 Length : ",len(num3))
	fmt.Println("Num 3 Capacity : ",cap(num3))

	fmt.Println("Let's copy num2 into num3 slice")
	copy(num3, num2)
	fmt.Println("Num 3 : ",num3)
	fmt.Println("Num 3 Length : ",len(num3))
	fmt.Println("Num 3 Capacity : ",cap(num3))


	fmt.Println("Let's check if both num2 and num3 slicer are equal : ")
	isSliceEqual := slices.Equal(num2, num3)
	fmt.Println("Is the slices Equal : ",isSliceEqual)
}


func mapPractice(){
	fmt.Println("Let's practice Map")
	var map1 = make(map[string]any)

	fmt.Println("Check the map 1 : ",map1)
	fmt.Println("Length of the map 1 : ", len(map1))

	fmt.Println("Let's add values to map 1 : ")
	map1["userId"] = 10
	map1["address"] = "Rampura"
	map1["role"] = 20
	map1["productId"] = 292552250
	map1["price"] = 4300
	fmt.Println("values of map1 : ",map1)

	var map2 = make(map[string]any)

	maps.Copy(map2, map1)
	fmt.Println("Copied Map1 on Map2 : ",map2)
	fmt.Println("Let's print a key that is not available : ",map2["destination"])

	fmt.Println("Let'c create a map in another way : ")

	map3 := map[string]any{"price":30,"ratings": 5}

	fmt.Println("Values of map3 : ", map3)
	fmt.Println("Single value of Map3 : ",map3["price"])
	
} 


func rangePractice(){
	mySlice := []string{"Imtiaz","chowdhury","mango","pinaple","fruits","Licy"}
	 fullLine := ""
	 
	// for i:=0; i<len(mySlice); i++{
	// 	fullLine = fullLine + mySlice[i]
	// }

	for index, data := range mySlice{
		fullLine = fullLine + data

		fmt.Println("Index : ",index)
		fmt.Println("Data : ",mySlice[index])
		fmt.Println("Charecters : ")
		for myKey, myValue := range data{
			fmt.Println(myKey,string(myValue))
		}
	}

	fmt.Println("Full Line : ",fullLine)
}

func functionPractice(num1,num2 int)(int,int){
	res := num1+num2
	resMulti2 := res*2
	return res, resMulti2
}

func printValueOfTheFunctionPractice(num1, num2 int){
	value1,value2 := functionPractice(num1, num2)
	fmt.Println("Value of the function : ",value1,value2)
}


func returnFunctionFromAFunction(val int) func(myName string) (string,int) {
	myNameValue := "Imtiaz"
	return func (myName string)(string,int){

		res := myNameValue + " Says hello to " + myName
		return res,val
	}
}

///Variadic Function besically means the user can insert as many input as much he want's to.
func variadicFucntionPractice(num ...int)int{
	total := 0
	for _,v:= range num{
		total = total + v
	}

	return total

}


///Struct Practice
type Order struct {
		id string
		amount float32
		status bool
		createdAt time.Time
	}
func structPractice ()Order{
	myOrder :=  Order{
		id: "1",
		amount: 30,
		status: true,
	}
return myOrder 
}


////Section : Pointer Practice
func pointerPractice(num1 *int, num2 *int) func(*int, *int) int {

	fmt.Println("Address of Outer num 1 : ",&num1)
	fmt.Println("Value At Address : ",*num1)
	fmt.Println()

	fmt.Println("Address of Outer num 2 : ",&num2)
	fmt.Println("Value At Address : ",*num2)
	fmt.Println()

	sumOfTwoDigit := func(num1 *int, num2 *int)int{

	fmt.Println("Address of Inner num 1 : ",&num1)
	fmt.Println("Value At Address : ",*num1)
	fmt.Println()

	fmt.Println("Address of Inner num 2 : ",&num2)
	fmt.Println("Value At Address : ",*num2)
	fmt.Println()

	res := *num1+*num2

	fmt.Println("Address of Inner num 1 : ",&num1)
	fmt.Println("Value At Address : ",*num1)
	fmt.Println()

	fmt.Println("Address of Inner num 2 : ",&num2)
	fmt.Println("Value At Address : ",*num2)


		return res 
	}
fmt.Println("---------------------------------------")
	fmt.Println("Value At Address : ",*num1)
	fmt.Println("Value At Address : ",*num2)
	fmt.Println("---------------------------------------")
	fmt.Println("")



	return sumOfTwoDigit

}



func dateTimePractice(){
	currentDateTime := time.Now();
	fmt.Println("Current Date And Time : ",currentDateTime)
	
	fmt.Println()
	fmt.Println("Let's show the Date in \"Day-Month-Year\" Format")
	fmt.Println()


	formatedDate := currentDateTime.Format("02-01-2006")
	fmt.Printf("Formated Date : %s",formatedDate)
	fmt.Println()


	fmt.Println()
	fmt.Println("Let's show the Time in \"Hour-Minute-Second\" Format")
	fmt.Println()
	


	formatedTimeAt24Hr := currentDateTime.Format("15-04-05")
	fmt.Printf("Formated Time in 24hr Format : %s",formatedTimeAt24Hr)
	fmt.Println()

	formatedTimeAt12Hr := currentDateTime.Format("03-04-05")
	fmt.Printf("Formated Time in 12Hr Format : %s",formatedTimeAt12Hr)
	fmt.Println()

	formatedTimeAt12HrAmPmFormat := currentDateTime.Format("03-04 PM")
	fmt.Printf("Formated Time in 12Hr AM/PM Format : %s",formatedTimeAt12HrAmPmFormat)
	fmt.Println()

	formatedTimeWithDayNameFormat := currentDateTime.Format("03-04 PM Monday")
	fmt.Printf("Formated Time in 12Hr AM/PM Format : %s",formatedTimeWithDayNameFormat)
	fmt.Println()

	dateAtDayMonthYearFormat := currentDateTime.Format("01-02-2006")

	layout_str := "01-02-2006"
	

	parseInTimeFormatFromDateFormat,_ := time.Parse(layout_str,dateAtDayMonthYearFormat)
	fmt.Printf("Formated Time in Time Format : %s",parseInTimeFormatFromDateFormat)
	fmt.Println()
	fmt.Printf("Formated Time in Formate Type : %T",parseInTimeFormatFromDateFormat)
	fmt.Println()
}