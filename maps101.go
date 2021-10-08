package main

import (
	"fmt"
)

func declareMap(){
	var mp map[string]string
	// mp["key"] = "value"
	fmt.Println(mp)
	mp2 := make(map[string]string)
	mp2["key"] = "value"
	fmt.Println(mp2)
	fmt.Println(mp2["key"])
}

func getKey(){
	mp := make(map[string]string)
	mp["key"] = "value"
	val, present := mp["key"]
	if present {		
		fmt.Println(val)
	}
	val2, present2 := mp["key2"]
	if present2 {		
		fmt.Println(val2)
	}
}

func deleteKey(){
	mp := make(map[string]string)
	mp["key"] = "value"
	fmt.Println(mp)
	delete(mp,"key")
	val, present := mp["key"]
	if present {		
		fmt.Println(val)
	}	
}

func mapOperations(){
	// Iterate over
	var mp =  map[string]string {
		"key":"value",
		"key2":"value2",
	}	
	for k,v := range(mp) {
		fmt.Println(k,v)
	}
}

func main(){
	fmt.Println("Maps 101")
	// declareMap()
	// getKey()
	// deleteKey()
	mapOperations()
}