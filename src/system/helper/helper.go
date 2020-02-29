package helper

import (
    "system/config"
    "net/http"
    "strings"
)

type Permission struct{}

func (param Permission) CheckDomainRequestPermission(request *http.Request) bool {

    // ambil acceptable request domain
    var clientRequestDomain = config.Server.AcceptableRequestDomain
    // ambil client host
    var clientHost = strings.Split(request.RemoteAddr, ":")[0]

    // jika tidak ada client host yang di
    // masukkan, berarti accept semua
    // host
    if len(clientRequestDomain) == 0 {
        return true
    }

    // cek apakah client request domain
    // sudah terdaftar
    for i := 0; i < len(clientRequestDomain); i++ {
        if clientRequestDomain[i] == clientHost {
            return true
        }
    }

    return false
}
