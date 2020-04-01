package main

import(
	"fmt"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Hello. This is the most easiest start you can have </h1>")
}
func main(){
	http.HandleFunc("/", HandlerFunc)
	http.ListenAndServe(":3000", nil)
}
