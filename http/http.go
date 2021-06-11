package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (f *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellow my handler")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "New Name"
	}
	fmt.Fprintf(w, "Hello World1 test %s", name)
}

func main() {

	//fmt.Println("Hello World")
	// 정적 처리
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello World")
		})

		http.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello World1")
		})

		http.HandleFunc("/test", testHandler)

		http.Handle("/my", &myHandler{})

		http.ListenAndServe(":3001", nil)
	*/

	//  먹스 활용 :동적
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World1")
	})

	mux.HandleFunc("/test", testHandler)

	mux.Handle("/my", &myHandler{})

	http.ListenAndServe(":3002", mux)

}
