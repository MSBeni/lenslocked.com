package main

import(
	"fmt"
	"net/http"
)

// Please check the https://github.com/gravityblast/fresh to install the fresh package for dynamic reloading and conf file
func HandlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Hello. This is the most easiest simple start you can have </h1>")
}
func main(){
	http.HandleFunc("/", HandlerFunc)
	http.ListenAndServe(":3000", nil)
}
