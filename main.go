package main

import (
	"fmt"
	"net/http"

	f "dockeriz/func"
)

func main() {
	// http.Handle("/func/", http.StripPrefix("/func/", http.FileServer(http.Dir("../func"))))
	http.HandleFunc("/styles/", f.ServeStyle)
	// http.HandleFunc("/styles/", Styles)
	fmt.Println("The server is working now :")
	http.HandleFunc("/", f.Welcom)
	http.HandleFunc("/ascii-art", f.Last)
	// starting the server at local host at port 8080
	fmt.Println("the server is running on localhost port 8082")
	fmt.Println("http://localhost:8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println(err)
	}
}
