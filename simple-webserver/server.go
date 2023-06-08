package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s received at /\n", r.Method)

	rb := ResponseBody{}
	s := http.StatusOK
	if r.Method != http.MethodGet {
		rb.Message = "Method not allowed!"
		s = http.StatusMethodNotAllowed
	} else {
		rb.Message = "Hello World!!"
		rb.Data = map[string]string{"foo": "Bar"}
	}

	rb.WriteJsonResponse(w, s)
}

func initServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	port := "8088"
	fmt.Println("Server started at http://localhost:" + port + ".")
	fmt.Print("Press ctrl + c to stop.\n\n")
	http.ListenAndServe(":"+port, mux)
}
