package loader

import (
    "html/template"
    "net/http"
    // "fmt"
)

var viewPath = "application/views/"
type Loader struct {}

func (param Loader) View(path string, response http.ResponseWriter, data interface{}) bool {
    var tmplt = template.Must(template.New(path).ParseFiles(viewPath + path))
    var error = tmplt.Execute(response, data)

    return error != nil
}
