package main

import (
	"fmt"
)

func main(){
	var1 := 65
	var2 := 85
	var3 := 35
	if (var1 > var2) && (var1 > var3){		
		fmt.Printf("%d is Greater than %d and %d",var1,var2,var3)		
	} else if (var2 > var3){
		fmt.Printf("%d is Greater than %d and %d",var2,var1,var3)		
	} else {
		fmt.Printf("%d is Greater than %d and %d",var3,var1,var2)		
	}
}