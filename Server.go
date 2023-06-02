package main

//
//import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"os"
//)
//
//var abobas = []Aboba{
//	{a: "a"},
//}
//
//func getRoot(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "This is my website!\n")
//}
//func getHello(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "Hello, HTTP!\n")
//}
//
//func postData(w http.ResponseWriter, r *http.Request) {
//	//reqBody, _ := ioutil.ReadAll(r.Body)
//	//groceries = nil
//	//var grocery Aboba
//	//json.Unmarshal(reqBody, &grocery)
//	//groceries = append(groceries, grocery)
//	//
//	//json.NewEncoder(w).Encode(groceries)
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Printf("server: could not read request body: %s\n", err)
//	}
//	err = json.Unmarshal(reqBody, abobas)
//	if err != nil {
//		panic(err)
//	}
//
//	p := abobas
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(p)
//
//}
//
//type Aboba struct {
//	a string
//}
//
//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", getRoot)
//	mux.HandleFunc("/hello", getHello)
//	mux.HandleFunc("/data", postData)
//	err := http.ListenAndServe("localhost:8080", mux)
//	if errors.Is(err, http.ErrServerClosed) {
//		fmt.Printf("server closed\n")
//	} else if err != nil {
//		fmt.Printf("error starting server: %s\n", err)
//		os.Exit(1)
//	}
//}
