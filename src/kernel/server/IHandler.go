package server

type IHandler interface {
	Build(server *Server)
}
