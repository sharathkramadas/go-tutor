package main 

import (
	"fmt"
)

func main(){
	var input string
	fmt.Println("What's your favourite car?")
	fmt.Scanf("%s",&input)
	fmt.Printf("Whoa! %s is a nice car!",input)
}