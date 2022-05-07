package user

type ServerOption func(s *Server)

func WebAddress(address string) ServerOption {
	return func(s *Server) {
		s.webAddress = address
	}
}

func AuthAddress(address string) ServerOption {
	return func(s *Server) {
		s.authAddress = address
	}
}
