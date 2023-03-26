package main 

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "http://example.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 
	}
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("Response status code: %s\n", resp.Status)
	fmt.Printf("Response body: %s\n", string(body))

}