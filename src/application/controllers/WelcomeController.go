package controllers

import (
    "system/loader"
    "net/http"
)

var load = new(loader.Loader)

type WelcomeController struct {}

func (param WelcomeController) Index(response http.ResponseWriter, request *http.Request){
    load.View("welcome.html", response, nil)
}
