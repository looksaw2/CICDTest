package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHelloWorld(t *testing.T){
	req ,_ := http.NewRequest("GET","/",nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloWorld)
	handler.ServeHTTP(rr,req)
	if status := rr.Code ; status != http.StatusOK {
		t.Errorf("You get %v but you want to get %v",rr.Code,http.StatusOK)
	} 
	expected := "Hello World"
	if s := rr.Body.String(); s != expected {
		t.Errorf("You get %v but you want to get %v",s,expected)
	}
}


func BenchmarkHelloWorld(b *testing.B) {
	req,_ := http.NewRequest("GET","/",nil)
	handler :=  http.HandlerFunc(HelloWorld)
	for i:= 0; i < b.N ; i++ {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr,req)
	}
}