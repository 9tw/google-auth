package main

import (
	"net/http"
	"oauth2/controller"
)

func main() {
	http.HandleFunc("/google/login", controller.GoogleLogin)
	http.HandleFunc("/google/callback", controller.GoogleCallback)

	http.ListenAndServe(":3000", nil)
}
