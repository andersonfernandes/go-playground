package handlers

import (
	"go-playground/simple-webserver/pkg/restserver"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	rb := restserver.ResponseBody{Message: "OK"}
	rb.WriteJsonResponse(w, http.StatusOK)
}
