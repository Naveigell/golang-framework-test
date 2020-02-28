package main

import (
    "system/config"
)

var server = new(config.Server)

func Init()  {
    server.HOST = "127.0.0.1"
    server.PORT = "4000"
}
