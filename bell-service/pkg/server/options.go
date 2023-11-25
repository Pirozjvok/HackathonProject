package server

import "net"

type Options func(s *Server)

func Port(addr string) Options {
	return func(s *Server) {
		s.httpServer.Addr = net.JoinHostPort("", addr)
	}
}
