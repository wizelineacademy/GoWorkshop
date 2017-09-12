package web

import (
	"encoding/json"
	"net/http"

	"github.com/gocraft/web"
	"github.com/gorilla/context"
)

// Context struct
type Context struct{}

// ListenAndServe func
func ListenAndServe() {
	ctx := new(Context)

	r := web.New(Context{}).
		Get("/", ctx.home)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)
	staticFolders := []string{"scripts", "styles"}
	for _, sf := range staticFolders {
		serveMux.Handle("/"+sf+"/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	}

	http.ListenAndServe(":8080", context.ClearHandler(serveMux))
}

// Ajax func
func (c *Context) Ajax(w web.ResponseWriter, r *web.Request, response interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resultJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
}
