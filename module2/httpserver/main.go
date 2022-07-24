package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request at %v\n", time.Now())

	// Check the required parameter name
	keys, parameterFound := r.URL.Query()["name"]
	if !parameterFound {
		w.WriteHeader(400)
		resp := make(map[string]string)
		resp["message"] = "missing required parameter [name]"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		return
	}

	// Set response header
	for k, v := range r.Header {
		fmt.Printf("%v: %v\n", k, v)
		w.Header().Add(k, v[0])
	}

	// Set message
	w.Header().Set("content-type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "hello, " + keys[0]
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func main() {
	// register router
	http.HandleFunc("/hello", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
