package restserver

import (
	"fmt"
	"net/http"
)

type Server struct {
	Mux  http.ServeMux
	Port string
}

func (s *Server) Start() {
	fmt.Println("Server started at http://localhost:" + s.Port + ".")
	fmt.Print("Press ctrl + c to stop.\n\n")
	http.ListenAndServe(":"+s.Port, &s.Mux)
}
