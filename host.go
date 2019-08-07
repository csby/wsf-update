package main

import (
	"fmt"
	"github.com/csby/wsf/types"
	"github.com/kardianos/service"
)

type Host struct {
	service service.Service
	program *Program
}

func (s *Host) ServiceName() string {
	return s.service.String()
}

func (s *Host) Interactive() bool {
	return service.Interactive()
}

func (s *Host) Run() error {
	if s.Interactive() {
		err := s.program.server.Run()
		if err != nil {
			return fmt.Errorf("run server error: %v", err)
		}
	} else {
		err := s.service.Run()
		if err != nil {
			return fmt.Errorf("run service error: %v", err)
		}
	}

	return nil
}

func (s *Host) Shutdown() error {
	return s.program.server.Close()
}

func (s *Host) Restart() error {
	return s.service.Restart()
}

func (s *Host) Start() error {
	return s.service.Start()
}

func (s *Host) Stop() error {
	return s.service.Stop()
}

func (s *Host) Install() error {
	return s.service.Install()
}

func (s *Host) Uninstall() error {
	err := s.service.Stop()
	if err != nil {
	}

	return s.service.Uninstall()
}

func (s *Host) Status() (types.ServerStatus, error) {
	status, err := s.service.Status()
	if err != nil {
		return types.ServerStatusUnknown, err
	}

	return types.ServerStatus(status), nil
}
