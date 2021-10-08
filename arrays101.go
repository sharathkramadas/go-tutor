package main

import (
	"fmt"
)

func declareArray(){
	var arr [4]int;
	fmt.Println(arr);
	arr[0] = 3
	fmt.Println(arr);
	arr[1] = 4
	arr[2] = 5
	arr[3] = 6
	fmt.Println(arr);
	fmt.Println(arr[3]);
}

func declareAndIntializeArray(){
	data := [5]string{"IronMan","Hulk","Thor","Hawkeye","Falcon"}
	fmt.Println(data);
}

func multiDatatypeArray(){	
	mda := []interface{}{"Area",5,3.14}
	fmt.Println(mda);
	fmt.Println(mda[2]);
}

func multiDimensionalArray(){
	matrix := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(matrix);
	fmt.Println(matrix[0][1]);
}

func arrayOperations(){
	arr := [3]int{1,2,3}
	fmt.Println("Length of an array is ",len(arr))
	fmt.Println("Iterate array data")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("First element of an array is ",arr[0])
	fmt.Println("Last element of an array is ",arr[len(arr)-1])
	arr2 := [3]int{1,2,3}
	fmt.Println("Does two arrays contain same values?",arr == arr2)
}

func main(){
	fmt.Println("Arrays 101")
	// Declare array
	declareArray()
	// Declare and initialize array
	declareAndIntializeArray()
	// Multi datatype array
	multiDatatypeArray()
	// Multi Dimensional Array
	multiDimensionalArray()
	// Array operations
	arrayOperations()
}