package main

import(
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Roadmap Project")

	fmt.Println("Server starting")
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Server running at port http://127.0.0.1:3000")
	http.ListenAndServe(":3000", nil)
}