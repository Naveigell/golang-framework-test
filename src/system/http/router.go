package http

import (
    "net/http"
    "fmt"
    // "net/url"
    "html/template"
    "log"
)

type function func(response http.ResponseWriter, request *http.Request)

type Router struct {}

type Method struct {
    Url string
    Method[] string
    Function[] function
    IsMainRouteActive bool
}

var methods[] Method

func (param Router) Get(url string, fn function)  {
    pass(url, http.MethodGet, fn)
}

func (param Router) Post(url string, fn function)  {
    pass(url, http.MethodPost, fn)
}

func pass(url string, m string, fn function)  {
    var method = Method {
        Url: url,
        Method: []string{m},
        Function: []function{fn},
        IsMainRouteActive: url == "/",
    }

    // cek semua daftar route
    for i := 0; i < len(methods); i++ {
        var temp = methods[i]
        // cek apakah urlnya sama
        if method.Url == temp.Url {
            // cek apakah ada route dengan method yang sama
            for index := 0; index < len(method.Method); index++ {
                if temp.Method[index] == m {
                    log.Fatal("There is route with the same METHOD in url --> ", method.Url)
                    return
                }
            }
            // jika sama, maka kita tidak perlu menambah ke
            // daftar route baru, cukup menambah ke dalam
            // array yang ada pada struct Method
            methods[i].Method = append(methods[i].Method, m)
            methods[i].Function = append(methods[i].Function, fn)
            return
        }
    }

    methods = append(methods, method)
}

func Serve(response http.ResponseWriter, request *http.Request){
    switch request.Method {
        case http.MethodGet:
            findRoute(response, request, http.MethodGet)
            break
        case http.MethodPost:
            findRoute(response, request, http.MethodPost)
            break
    }
}

func (param Router) Save() {
    var findActiveHome = false
    // looping masing - masing methods
    for i := 0; i < len(methods); i++ {
        var method = methods[i]

        // cek apakah ada main route
        if !findActiveHome {
            findActiveHome = method.Url == "/"
        }

        http.HandleFunc(method.Url, Serve)
    }

    if !findActiveHome {
        // jika tidak ada main route maka
        // buat main route dengan menonaktifkan
        // IsMainRouteActive
        var method = Method {
            Url: "/",
            Method: []string{http.MethodGet},
            IsMainRouteActive: false,
        }
        methods = append(methods, method)

        http.HandleFunc("/", Serve)
    }
}

func findRoute(response http.ResponseWriter, request *http.Request, m string)  {
    // ambil url path saat ini
    var urlPath = request.URL.Path
    var isFound = false

    // looping ke semua route yang didaftarkan
    // dengan url
    for index := 0; index < len(methods); index++ {
        var method = methods[index]
        // cek apakah main routenya active dan urlnya = "/"
        if !method.IsMainRouteActive && method.Url == "/" {
            break
        }

        if method.Url == urlPath {
            // looping masing - masing method yang dimiloki
            // oleh var method
            for i := 0; i < len(method.Method); i++ {
                if method.Method[i] == m {
                    // jika sama maka panggil fungsinya
                    method.Function[i](response, request)
                    // jadikan found = true lalu break
                    isFound = true
                    break
                }
            }
        }
    }

    if !isFound {
        var tmplt = template.Must(template.New("404.html").ParseFiles("application/views/error/404.html"))
        var error = tmplt.Execute(response, nil)
        if error != nil {
            fmt.Println("Error when serving 404 file")
        }
    }
}
