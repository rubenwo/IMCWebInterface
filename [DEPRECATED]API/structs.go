package API

type Server struct {
	Port    string
	Running bool
}

type response struct {
	Data string `json:"data,omitempty"`
}

type Credentials struct {
	Username string `json:"username, omitempty"`
	Password string `json:"password, omitempty"`
}

type loginResponse struct {
	Valid bool   `json:"valid,omitempty"`
	UID   string `json:"uid,omitempty"`
}
