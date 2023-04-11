// Реализовать паттерн «адаптер» на любом примере.
package main

type Inter interface {
	Request()
}

type Server struct {
	token string
}

type AdapterServer struct {
	*Server
}

func (s *Server) ServerRequest() {
	s.token = "qwerty-asdfgh-zxcvbn"
}

func (a *AdapterServer) Request() {
	a.ServerRequest()
}

func main() {

}
