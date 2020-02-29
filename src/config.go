package main

import (
    "system/config"
)

func Init()  {
    config.Server.HOST = "127.0.0.1"
    config.Server.PORT = "4000"

    config.Server.AcceptRequestDomain = []string {
        "127.0.0.1",
    }
}
