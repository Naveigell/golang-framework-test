package loader

import (
    "html/template"
    "net/http"
    "strings"
)

var viewPath = "application/views/"
type Loader struct {}

func (param Loader) View(path string, response http.ResponseWriter, data interface{}) bool {
    var splitedPathString = strings.Split(path, "/")

    var tmplt = template.Must(template.New(splitedPathString[len(splitedPathString) - 1]).ParseFiles(viewPath + path))
    var error = tmplt.Execute(response, data)

    return  error != nil
}
