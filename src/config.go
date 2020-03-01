package main

import (
    "system/config"
)

func Init()  {
    //
    // SERVER
    //

    config.Server.HOST = "127.0.0.1"
    config.Server.PORT = "4000"

    config.Server.AcceptableHostRequest = []string {
        "http://" + config.Server.HOST + ":" + config.Server.PORT,
    }

    //
    // CORS
    //

    config.Cors.GlobalCorsEnable = true
}
