package user_v1

import "net/http"

const PREFIX string = "/user"

func HandleRequest(mux *http.ServeMux) {
	mux.HandleFunc("POST "+PREFIX, create)
}

func create(w http.ResponseWriter, r *http.Request) {
	
}
