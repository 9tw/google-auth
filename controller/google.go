package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"oauth2/config"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Fprintln(res, "states dont match")
		return
	}

	code := req.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintln(res, "Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Fprintln(res, "User Data fetch Failed")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "Json Parsing Failed")
	}

	fmt.Fprintln(res, string(userData))
}
