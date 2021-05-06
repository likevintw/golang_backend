package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"http_template/public"
)

type BackendTester struct {
	url string
}

func (self BackendTester) UserHandler(
	dataMap public.ClientInfomationFormat) public.BackendInformation {

	// send http request
	client := &http.Client{}
	dataByte, _ := json.Marshal(dataMap)
	request, err := http.NewRequest(http.MethodGet,
		self.url, bytes.NewBuffer(dataByte))
	if err != nil {
		fmt.Println("http.NewRequest fail")
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do(request) fail")
		panic(err)
	}

	// Get Response with Byte
	backendResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println(string(backendResponse))
	var responseJson public.BackendInformation

	// Conver Response Byte to Json
	err = json.Unmarshal(backendResponse, &responseJson)
	if err != nil {
		fmt.Println(("%v"), err)
	}
	fmt.Println("responseJson.Status: ", responseJson.Status)
	fmt.Println("responseJson.Message: ", responseJson.Message)

	return responseJson
	// return string(backendResponse) // json Byte type
}

func main() {
	// define
	url := "http://0.0.0.0:5000/user"
	tester := BackendTester{url}
	var dataMap public.ClientInfomationFormat
	// send a request
	dataMap.Behavior = "login"
	dataMap.Username = "ken"
	dataMap.Password = "111111"
	responseMap := tester.UserHandler(dataMap)
	fmt.Println("Login Test:", responseMap, reflect.TypeOf(responseMap))
}
