package main

import (
    "net/http"
    "log"
)

func RunApp()  {

    Init()
    HandleRoutes()

    log.Println("Listen on port : ", server.PORT)
    log.Fatal(http.ListenAndServe(server.HOST + ":" + server.PORT, nil))
}
