package main

import (
    ht "system/http"
    "net/http"
    // "fmt"
    "system/loader"
    // "html/template"
    // "net/url"
)

var routes = new(ht.Router)
var load = new(loader.Loader)

func test (response http.ResponseWriter, request *http.Request)  {
    // var query = request.URL.Query()
    // fmt.Println("/login --> GET", query)

    var error = load.View("welcome.html", response, nil)
    if error {

    }

    // var tmplt = template.Must(template.New("welcome.html").ParseFiles("application/views/welcome.html"))
    // var tmpltError = tmplt.Execute(response, nil)
    // if tmpltError != nil {
    //     http.Error(response, "Template not reloading", http.StatusInternalServerError)
    // }
}

func HandleRoutes()  {
    routes.Get("/login", test);
    routes.Post("/login", func (response http.ResponseWriter, request *http.Request)  {
        // fmt.Println("/login --> POST")
    })
}
