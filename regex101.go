package main

import (
	"fmt"
	"regexp"
)

func main(){
	str := `
	NAME="Alpine Linux"
ID=alpine
VERSION_ID=3.14.2
PRETTY_NAME="Alpine Linux v3.14"
HOME_URL="https://alpinelinux.org/"
BUG_REPORT_URL="https://bugs.alpinelinux.org/"
	`
	r := regexp.MustCompile("\nID=(.*)")
	osFlavour := r.FindStringSubmatch(str)[1]
	fmt.Println(osFlavour)	
}