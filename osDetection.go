package main

import (
	"fmt"
	"runtime"
	"strings"
	"io/ioutil"	
	"regexp"	
)

func getLinuxFlavour(content string) string{
	r := regexp.MustCompile("\nID=(.*)")
	flavour := r.FindStringSubmatch(content)[1]
	return strings.Replace(flavour,"\"","",-1)
}

func main(){	
	if runtime.GOOS == "linux" {		
		data, err := ioutil.ReadFile("/etc/os-release")
		if err != nil {
			fmt.Println("You are running Linux OS, But we could not detect linux flavour!")
		} else {		
			content := string(data)
			linuxFlavour := getLinuxFlavour(content)			
			fmt.Printf("You are running %s OS",strings.Title(linuxFlavour))
		}
	} else {
		fmt.Printf("You are running %s OS",strings.Title(runtime.GOOS))
	}
}