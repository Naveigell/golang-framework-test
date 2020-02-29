package main

import (
    "net/http"
    "log"
    "system/config"
)

func RunApp()  {

    Init()
    HandleRoutes()

    log.Println("Listen on port : ", config.Server.PORT)
    log.Fatal(http.ListenAndServe(config.Server.HOST + ":" + config.Server.PORT, nil))
}
