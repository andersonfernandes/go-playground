package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s received at /\n", r.Method)

	rb := &ResponseBody{
		Message: "Hello World!!",
		Data: map[string]string{
			"foo": "Bar",
		},
	}

	rb.WriteJsonResponse(w)
}

func initServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	port := "8088"
	fmt.Println("Server started at http://localhost:" + port + ".")
	fmt.Print("Press ctrl + c to stop.\n\n")
	http.ListenAndServe(":"+port, mux)
}
