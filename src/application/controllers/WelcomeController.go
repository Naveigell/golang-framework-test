package controllers

import (
    "system/loader"
    "system/helper"
    "net/http"
)

var load = new(loader.Loader)
var permission = new(helper.Permission)

type WelcomeController struct {}

func (param WelcomeController) Index(response http.ResponseWriter, request *http.Request){
    load.View("welcome.html", response, nil)
}
