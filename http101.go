package main

import (
	"fmt"
	"net/http"
	"io"
)

func httpGet(){
	res, err := http.Get("https://we45.com")
	if err != nil {
		fmt.Println("Could not get this page!")
	} else{
		if res.StatusCode == 200 {
			defer res.Body.Close()
			content, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Could not read response!")
			} else {
				fmt.Println(string(content))
				fmt.Println(res.Header["Content-Type"][0])
			}
		}
		
	}
}

func main(){	
	httpGet()
}