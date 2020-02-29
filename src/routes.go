package main

import (
    http "system/http"
    "system/loader"
    "application/controllers"
)

var routes = new(http.Router)
var load = new(loader.Loader)

var welcome = new(controllers.WelcomeController)

func HandleRoutes()  {
    routes.Get("/", welcome.Index)

    routes.Save()
}
