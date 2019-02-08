package FrontEnd

import (
	"fmt"
	"html/template"
	"net/http"
)

type Server struct {
	Port string
}

type rootPage struct {
	Title string
}

func (s *Server) Start() {
	if s.Port == "" {
		s.Port = "80"
	}
	//http.HandleFunc("/", rootHandler)
	fs := http.FileServer(http.Dir("FrontEnd/public"))
	http.Handle("/", fs)
	http.ListenAndServe(":"+s.Port, nil)
}
func rootHandler(w http.ResponseWriter, req *http.Request) {
	p := rootPage{Title: "IMC Configurator"}
	t, _ := template.ParseFiles("FrontEnd/template.html")
	fmt.Println(t.Execute(w, p))
}
