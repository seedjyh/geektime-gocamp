package admin

type ServerOption func(s *Server)

func WebAddress(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

func DataSourceName(dsn string) ServerOption {
	return func(s *Server) {
		s.dataSourceName = dsn
	}
}
