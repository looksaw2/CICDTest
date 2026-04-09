package main

import (
	"fmt"
	"net/http"
)

var PORT = 8081
func HelloWorld(w http.ResponseWriter ,r *http.Request){
	w.Write([]byte("Hello World"))
}

func main(){
	http.HandleFunc("/",HelloWorld)
	fmt.Printf("Server start at %d",PORT)
	http.ListenAndServe(":8081",nil)
}