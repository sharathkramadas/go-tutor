package main

import (
	"fmt"
	"runtime"
	"strings"
	"io/ioutil"	
	"os/exec"
	"regexp"	
)

func listAlpinePackages(){
	cmd := "apk info -v"
	out, err := exec.Command("sh","-c",cmd).Output()
	if err != nil {
		fmt.Println("Could not run this command!")
	} else {
		packages := strings.Split(string(out),"\n")
		if packages != nil {						
			var i int;
			for i = 0; i < len(packages); i++ {		
				packageSplit := strings.Split(packages[i],"-")
				if len(packageSplit) > 2 {								
					packageName := strings.Join(packageSplit[:len(packageSplit)-2],"-")
					packageVersion := packageSplit[len(packageSplit)-2] +"-"+ packageSplit[len(packageSplit)-1]
					fmt.Println(packageName,packageVersion)
				}
			}					
		}
	}
}

func listUbuntuPackages(){
	cmd := "apt list --installed"
	out, err := exec.Command("sh","-c",cmd).Output()
	if err != nil {
		fmt.Println("Could not run this command!")
	} else {
		packages := strings.Split(string(out),"\n")
		if packages != nil {						
			var i int;
			for i = 0; i < len(packages); i++ {	
				packageSplit := strings.Split(packages[i]," ")
				if len(packageSplit) > 2 {								
					packageVersion := packageSplit[1]
					packageName := strings.Split(packageSplit[0],"/")[0]				
					fmt.Println(packageName,packageVersion)
				}
			}					
		}
	}
}

func listCentosPackages(){
	cmd := "yum list installed"
	out, err := exec.Command("sh","-c",cmd).Output()
	if err != nil {
		fmt.Println("Could not run this command!")
	} else {
		packages := strings.Split(string(out),"\n")
		if packages != nil {						
			var i int;
			for i = 0; i < len(packages); i++ {					
				packageSplit := strings.Fields(packages[i])
				if len(packageSplit) > 2 {					
					packageName := packageSplit[0]
					packageVersion := packageSplit[1]
					fmt.Println(packageName,packageVersion)
				}				
			}					
		}
	}
}

func getLinuxFlavour(content string) string{
	r := regexp.MustCompile("\nID=(.*)")
	flavour := r.FindStringSubmatch(content)[1]
	return strings.Replace(flavour,"\"","",-1)
}

func main(){	
	if runtime.GOOS == "linux" {		
		data, err := ioutil.ReadFile("/etc/os-release")
		if err != nil {
			fmt.Println("Could not open file!")
		} else {		
			content := string(data)
			linuxFlavour := getLinuxFlavour(content)			
			switch linuxFlavour {
				case "alpine": listAlpinePackages()
				case "ubuntu": listUbuntuPackages()
				case "debian": listUbuntuPackages()
				case "centos": listCentosPackages()
				case "fedora": listCentosPackages()
			default:
				fmt.Println("Could not find OS flavour!")
			}			
		}
	} else {
		fmt.Printf("Not supporting %s OS, For now!",strings.Title(runtime.GOOS))
	}
}