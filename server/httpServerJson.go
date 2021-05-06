package server

import (
	"encoding/json"
	"fmt"
	"http_template/public"
	"io/ioutil"
	"net/http"
)

type HttpServerJson struct {
}

func (handler *HttpServerJson) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// define retrun setting
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var response public.BackendInformation

	// Get Json Data From Clien
	r.ParseForm()
	dataString, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Status = "500"
		response.Message = "ioutil.ReadAll(r.Body) fail"
		json.NewEncoder(w).Encode(response)
		return
	}

	var dataJson public.ClientInfomationFormat
	json.Unmarshal(dataString, &dataJson) // type byte to json
	fmt.Println(dataJson)
	fmt.Println(dataJson.Behavior)
	fmt.Println(dataJson.Username)
	fmt.Println(dataJson.Password)
	fmt.Println(dataJson.Token)
	fmt.Println()

	// Set Response
	response.Status = "200"
	response.Message = "ok"
	json.NewEncoder(w).Encode(response) // return Map Type
	return
}
