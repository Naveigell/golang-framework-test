package helper

import (
    "system/config"
    "net/http"
)

type Permission struct{}

func enableCors(response http.ResponseWriter, domain string)  {
    response.Header().Set("Access-Control-Allow-Origin", domain)
    response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Origin, Authorization")
}

func (param Permission) EnableCors(response http.ResponseWriter, request *http.Request) bool {

    // ambil acceptable request domain
    var clientRequestDomain = config.Server.AcceptableHostRequest
    // ambil client host
    var clientOrigin = request.Header.Get("Origin")

    // jika tidak ada client host yang di
    // masukkan, berarti accept semua
    // origin
    if len(clientRequestDomain) == 0 {
        enableCors(response, "*")
        return false
    }

    // cek apakah client request domain
    // sudah terdaftar
    for i := 0; i < len(clientRequestDomain); i++ {
        if clientRequestDomain[i] == clientOrigin {
            enableCors(response, clientOrigin)

            return false
        }
    }

    return true
}
