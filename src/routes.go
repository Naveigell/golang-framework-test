package main

import (
    http "system/http"
    "application/controllers"
)

var routes = new(http.Router)

var welcome = new(controllers.WelcomeController)
func HandleRoutes()  {
    routes.Get("/api/v1/news-app", welcome.Index)

    routes.Save()
}
