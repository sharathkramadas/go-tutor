package main

import (
	"fmt"
)

type Class struct{
	object string
}

func (c Class) methodWithoutArgumentsAndReturnValue() {
	fmt.Println("Calling methodWithoutArgumentsAndReturnValue!")
}

func (c Class) methodWithReturnValuesAndNoArguments() string{
	return "Calling methodWithReturnValuesAndNoArguments!"
}

func (c Class) methodWithReturnValuesAndWithArguments(value string) string{
	return value
}

func (c Class) methodWithoutReturnValuesAndWithArguments(value string){
	fmt.Println(value)
}

func (c Class) accessingClassObjects(){
	fmt.Println(c.object)
}

func main(){
	fmt.Println("Struct as Class!")
	c := Class{object: "Calling Object!"}
	fmt.Println(c.object)
	c.object = "Calling Modified Object!"
	c.accessingClassObjects()
	fmt.Println(c.methodWithReturnValuesAndNoArguments())
	fmt.Println(c.methodWithReturnValuesAndWithArguments("Calling methodWithReturnValuesAndWithArguments!"))
	c.methodWithoutReturnValuesAndWithArguments("Calling methodWithoutReturnValuesAndWithArguments!")
	c.methodWithoutArgumentsAndReturnValue()
}

