package http

import (
    "net/http"
    "fmt"
    // "net/url"
    "log"
    "html/template"
)

type function func(response http.ResponseWriter, request *http.Request)

type Router struct {}

type Method struct {
    Url string
    Method[] string
    Function[] function
}

var methods[] Method

func (param Router) Get(url string, fn function)  {
    // pass url, method dan function untuk kemudian di masukkan ke daftar route
    pass(url, http.MethodGet, fn)
}

func (param Router) Post(url string, fn function)  {
    // pass url, method dan function untuk kemudian di masukkan ke daftar route
    pass(url, http.MethodPost, fn)
}

func pass(url string, m string, fn function)  {
    // buat variable sementara untuk menampung nilai - nilai yang ada di struct Method
    var method = Method {
        Url: url,
        Method: []string{m},
        Function: []function{fn},
    }

    // cek semua daftar route
    for i := 0; i < len(methods); i++ {
        var temp = methods[i]
        // cek apakah urlnya sama
        if method.Url == temp.Url {
            // cek apakah ada route dengan method yang sama
            for index := 0; index < len(method.Method); index++ {
                if temp.Method[i] == m {
                    log.Fatal("There is route with the same METHOD")
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

    // jika url tidak terdaftar, maka daftarkan
    methods = append(methods, method)
    http.HandleFunc(method.Url, Serve)
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

func findRoute(response http.ResponseWriter, request *http.Request, m string)  {
    // ambil url path saat ini
    var urlPath = request.URL.Path
    var isFound = false

    // looping ke semua route yang didaftarkan
    // dengan url
    for index := 0; index < len(methods); index++ {
        var method = methods[index]

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
        // jika daftar route tidak ditemukan maka tampilkan 404 file note found
        var tmplt = template.Must(template.New("404.html").ParseFiles("application/views/error/404.html"))
        var error = tmplt.Execute(response, nil)
        if error != nil {
            fmt.Println("Error when serving file")
        }
        fmt.Println("Route not found!")
    }
}
