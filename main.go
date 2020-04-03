package main

import (
	"fmt"
	"net/http"
)
//#########################################BASIC GO SERVER FILE#########################################################
// Please check the https://github.com/gravityblast/fresh to install the fresh package for dynamic reloading and conf file
func HandlerFunc(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/"{
		fmt.Fprint(w,"<h1>Hello. This is the most easiest simple start you can have </h1>")
	}else if r.URL.Path == "/contact"{
		fmt.Fprint(w,"To get in touch, please send an email to <a href \"mailto:support@lenslocked.com\">support@lenslocked.com</a>.")
	}

}
func main(){
	http.HandleFunc("/", HandlerFunc)
	http.ListenAndServe(":3000", nil)
}
//######################################################################################################################