package main

import (
	"fmt"
)

func declareSlice(){
	var sl []int;
	fmt.Println(sl);	
}

func declareAndIntializeSlice(){
	sl := []int{1,2,3,4,5};
	fmt.Println(sl);	
}

func multiDatatypeSlice(){	
	sl := []interface{}{"String",1,3.14}
	fmt.Println(sl);
}

func sliceOperations(){
	var sl []int;
	// Append to slice
	sl = append(sl,1)
	fmt.Println(sl);
	sl = append(sl,2,3,4)
	fmt.Println(sl);
	// Append to Multi datatype slice
	var mds []interface{}
	mds = append(mds,1)
	mds = append(mds,"String")
	mds = append(mds,3.14)
	fmt.Println(mds)
	fmt.Println(mds[0:2])
	fmt.Println(mds[1:2])
	// Iterate over slice
	for i := 0; i < len(mds); i++ {
		fmt.Println(mds[i])
	}
	for k,v := range mds {
		fmt.Println(k,v)
	}
}

func main(){
	fmt.Println("Slices 101")
	// Declare Slice
	declareSlice()
	// Declare and initialize Slice
	declareAndIntializeSlice()
	// Multi datatype Slice
	multiDatatypeSlice()	
	// Slice operations
	sliceOperations()
}