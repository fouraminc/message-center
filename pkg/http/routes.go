package http

func (s *server) routes() {

	s.router.GET("/", s.handleRoot())
	s.router.GET("/ping", s.handlePing())
}
