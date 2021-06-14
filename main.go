package main

import (
	"net/http"

	//"myapp"
	"github.com/yclee7/gowork/myapp"
)

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

	http.ListenAndServe(":3002", myapp.NewHttpHandler())

}
