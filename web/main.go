package main

import (
	"flag"
	"net/http"
	"log"
)

func main(){
	var addr = flag.String("addr", ":8081", "webサイトのアドレス")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir("public"))))
	log.Println("webサイトのアドレス:", *addr)
	http.ListenAndServe(*addr, mux)
}




