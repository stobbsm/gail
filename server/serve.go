package server

import "crypto/tls"

type Server struct {
	ListenAddr string
	Hostname   string
	listener   tls.Conn
	opts       []Option
}

type Option interface {
	Name() string
	Value() string
}

func (s *Server) Listen(listenAddr, opts ...Option) (*Server, error) {
	return nil, nil
}
