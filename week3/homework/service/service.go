package service

import (
	"context"
	"net/http"
)

type handler struct {
	content string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte(h.content))
}

type service struct {
	s *http.Server
}

func New(addr string, content string) *service {
	return &service{
		&http.Server{
			Addr:    addr,
			Handler: &handler{content: content},
		},
	}
}

func (s *service) Run() error {
	return s.s.ListenAndServe()
}

func (s *service) Stop() {
	_ = s.s.Shutdown(context.Background())
}
