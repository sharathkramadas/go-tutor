package main

import (
	"fmt"
	"os"
	"encoding/json"			
)

// Write JSON to a file

type Person struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserInfo struct {
    FirstName string
    LastName string
}

type User struct {
    Username string 
    Email string  
    Info  UserInfo 
}

func structToJSON(){
	i := UserInfo{"John","Doe"}
    u := User{"John", "john@mail.com", i}
    data, err := json.Marshal(u)
    if err != nil {
        fmt.Println(err)
    }    
    fmt.Println(string(data))
}


func writeJSONTofile(){
	file, err := os.Create("test.json")
	if err != nil {
		fmt.Println("Could not create a JSON file!")
	}
	if file != nil {		
		d := Person{Name:"John", Email:"john@doe.com"}
		jsonData, err := json.Marshal(d)
		fmt.Println(string(jsonData))
		if err == nil {				
			file.WriteString(string(jsonData))
			fmt.Println("JSON data saved successfully!")
		}
	}
	file.Close()
}

func stringToJSON(){
	s := `{"name":"John","email":"john@doe.com"}`
	p := Person{}
	json.Unmarshal([] byte(s), &p)
	fmt.Println(p)
	fmt.Println(p.Name)
}

func main(){
	fmt.Println("JSON Operations 101")	
	structToJSON()
	// writeJSONTofile()
	// stringToJSON()
}