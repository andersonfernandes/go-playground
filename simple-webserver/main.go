package main

import (
	"go-playground/simple-webserver/pkg/restserever"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s received at /\n", r.Method)

	rb := restserever.ResponseBody{}
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

func main() {
  server := restserever.Server{
    Mux: *http.NewServeMux(),
    Port: "8088",
  }

  server.Mux.HandleFunc("/", rootHandler)
  
  server.Start()
}
