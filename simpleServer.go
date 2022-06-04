package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", simpleServer)
	http.ListenAndServe(readfile(), nil)
}

func simpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func readfile() string {
	b, err := ioutil.ReadFile("app.properties")
	if err != nil {
		panic(err)
	}
	details := make(map[string]string)
	json.Unmarshal([]byte(b), &details)
	return ":" + details["port"]
}
