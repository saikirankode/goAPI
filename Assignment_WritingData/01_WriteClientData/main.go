package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func main() {
	client := &http.Client{}

	color := "black"
	breed := "bulldog"

	//url := fmt.Sprintf("http://localhost:8087/dogs?color=%v&breed=%v", color, breed)
	//url := fmt.Sprintf("http://localhost:8100/readFrmHeader?color=%v&breed=%v", color, breed)
	url := fmt.Sprintf("http://localhost:8099/readdogs?color=%v&breed=%v", color, breed)
	reqBody := []byte("hello")

	//req, _ := http.NewRequest("GET", url, nil)
	//req, _ := http.NewRequest("POST", url, nil)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	//req.Header.Set("color", "green")
	//req.Header.Set("breed", "pup")

	fmt.Println(req)
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured while sending request to the server")
	}
	defer res.Body.Close()
	res_Body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res.Status)
	fmt.Println(string(res_Body))

}
