package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func post() {

	jsonData := map[string]string{"Firstname": "Himanshu", "Lastname": "Phoenix"}
	jsonValue, err := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", "http://httpbin.org/post", bytes.NewBuffer(jsonValue))
	//content type could be application/json OR x-www-form-urlencoded
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	response, err := client.Do(request)
	//OR//
	//response,err := http.Post("http://httpbin.org/post","application/json",bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("http reques failed with => %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func put() {

	todo := Todo{1, 2, "Good Morning Heman", true}
	jsonreq, err := json.Marshal(todo)
	req, _ := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonreq))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	clientput := &http.Client{}
	respo, errr := clientput.Do(req)
	if err != nil {
		log.Fatalln(errr)
	}
	defer respo.Body.Close()
	bodyBytes, err := ioutil.ReadAll(respo.Body)
	if err != nil {
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		var todoStruct Todo
		json.Unmarshal(bodyBytes, &todoStruct)
		fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
	}
}

func delete() {
	todo := Todo{1, 2, "Good Morning Heman", true}
	jsonreq, err := json.Marshal(todo)
	req, _ := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonreq))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	clientput := &http.Client{}
	respo, errr := clientput.Do(req)
	if err != nil {
		log.Fatalln(errr)
	}
	defer respo.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(respo.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)

}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Printf("http reques failed with => %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}

	//POST API
	post()
	//PUT API
	put()

	//delete API
	delete()

}
