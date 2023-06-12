package handlers

import (
	"go-playground/simple-webserver/pkg/restserever"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	rb := restserever.ResponseBody{Message: "OK"}
	rb.WriteJsonResponse(w, http.StatusOK)
}
