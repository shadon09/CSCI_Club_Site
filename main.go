
package myapp


import (
	"html/template"
	"net/http"
)

var tpl* template.Template

func init(){

	tpl = template.Must(template.ParseGlob("templates/*.html"))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))



	http.HandleFunc("/", index)
}

func index(res http.ResponseWriter, req* http.Request){
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	//so when user enters main webpage we make a cookie. ok
	tpl.ExecuteTemplate(res, "index.html", nil)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}
