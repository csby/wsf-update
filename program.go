package main

import (
	"fmt"
	"github.com/kardianos/service"
)

type Program struct {
	server *Server
}

func (s *Program) Close() error {
	if s.server != nil {
		return s.server.Close()
	}
	return nil
}

func (s *Program) Start(svc service.Service) error {
	fmt.Println("info, service '", svc.String(), "' started")
	go s.run()
	return nil
}

func (s *Program) Stop(svc service.Service) error {
	fmt.Println("info, service '", svc.String(), "' stopped")

	if s.server != nil {
		s.server.Close()
	}

	return nil
}

func (s *Program) run() {
	if s.server != nil {
		err := s.server.Run()
		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}
