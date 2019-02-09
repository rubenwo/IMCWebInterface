package FrontEnd

import (
	"net/http"
)

type Server struct {
	Port string
}

var fs = http.FileServer(http.Dir("FrontEnd/public"))

func (s *Server) Start() {
	if s.Port == "" {
		s.Port = "80"
	}
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":"+s.Port, nil)
}
func rootHandler(w http.ResponseWriter, req *http.Request) {
	fs.ServeHTTP(w, req)
}
