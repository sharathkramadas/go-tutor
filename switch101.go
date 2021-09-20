package main

import (
	"fmt"
)

func main(){
	var option int64
	fmt.Println("Select your option \n1.First\n2.Second\n3.Third")
	fmt.Scanf("%d",&option)		
	switch option {
		case 1: 
			fmt.Printf("You have selected the option %d",option)
		case 2: 
			fmt.Printf("You have selected the option %d",option)
		case 3: 
			fmt.Printf("You have selected the option %d",option)
		default:
			fmt.Println("Wrong option selection!")
	}
}