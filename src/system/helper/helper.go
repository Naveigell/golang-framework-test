package helper

import (
    "system/config"
    "net/http"
    "strings"
)

type Permission struct{}

func (param Permission) CheckDomainRequestPermission(request *http.Request) bool {

    var clientRequestDomain = config.Server.AcceptRequestDomain
    var clientHost = strings.Split(request.RemoteAddr, ":")[0]

    if len(clientRequestDomain) == 0 {
        return true
    }

    for i := 0; i < len(clientRequestDomain); i++ {
        if clientRequestDomain[i] == clientHost {
            return true
        }
    }

    return false
}
