package main

import (
	"fmt"
	"os"	
	"io/ioutil"
)

// Create a file
// Write to a file
// Write Image to a file
// Append to a file
// Read from a file
// Delete a file


func createFile(){
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Could not create a file!")
	}
	if file != nil {
		fmt.Printf("File %s created successfully!\n", file.Name())
	}
	defer file.Close()
}


func writeToFile(){
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Could not create a file!")
	}
	if file != nil {
		fmt.Println("Added data successfully!")
	}
	file.WriteString("Test data")
	defer file.Close()
}


func appendToFile(){
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)	
	if err != nil {
		fmt.Println("Error appending to file!")
	}
	if file != nil {
		file.WriteString("\nAdding some more data!")
	} 
	defer file.Close()
}


func readFromFile(){
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Could not open a file!")
	} else {
		fmt.Println(string(data))
	}	
}


func deleteFile(){
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println("File could not be found!")
	} else {
		fmt.Println("File deleted successfully!")
	}
}


func main(){
	fmt.Println("File Operations 101")
	createFile()
	writeToFile()	
	readFromFile()
	appendToFile()
	readFromFile()
	deleteFile()
}