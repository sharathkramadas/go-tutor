package main

import (
	"fmt"
	"strconv"
	"strings"
)

func declareVariable(){
	var s1 string
	s1 = "String Ops!"
	fmt.Println(s1)
}

func lengthOfString(){
	s1 := "IronMan"
	fmt.Printf("Length of string `%s` is %d\n",s1,len(s1))
}

func integerConvertToString(){
	i := 1
	fmt.Println("Converted string is ",strconv.Itoa(i))
}

func floatConvertToString(){
	f := 1.0
	fmt.Println("Converted string is ",strconv.FormatFloat(f, 'E', -1, 32))
}

func utilityFunctions(){
	s1 := "wanda vision"
	fmt.Println(strings.Title(s1))	
	fmt.Println(strings.ToUpper(s1))
	s2 := "THOR"
	fmt.Println(strings.ToLower(s2))
	s3 := "antman"
	fmt.Println(strings.Replace(s3,"ant","bat",-1))
	s4 := "Super"
	s5 := "Girl"
	fmt.Println(s4+s5)
	fmt.Printf("%s %s\n",s4,s5)
	s6 := "Spider Man"
	fmt.Println(strings.Split(s6," "))
	s7 := []string{"Black","Widow"}	
	fmt.Println(strings.Join(s7," "))
}

func main(){
	declareVariable()
	lengthOfString()
	integerConvertToString()
	floatConvertToString()	
	utilityFunctions()
}