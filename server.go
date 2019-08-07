package main

import (
	"encoding/json"
	"fmt"
	"github.com/csby/wsf/types"
	"github.com/kardianos/service"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	path       string
	httpServer *http.Server
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 9606)
	fmt.Println("http server running on \"", addr, "\"")
	s.httpServer = &http.Server{Addr: addr, Handler: s}
	err := s.httpServer.ListenAndServe()
	s.httpServer = nil
	return err
}

func (s *Server) Close() error {
	if s.httpServer != nil {
		return s.httpServer.Close()
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := &Result{}
	if r.Method != "POST" {
		result.Code = 1
		result.Error = fmt.Sprintf("method '%s' not support", r.Method)
		result.Output(w)
		return
	}

	argument := &Argument{}
	err := json.NewDecoder(r.Body).Decode(argument)
	if err != nil {
		result.Code = 2
		result.Error = fmt.Sprintf("input error: %v", err)
		result.Output(w)
		return
	}

	if argument.Action == "info" {
		result.Name = moduleName
		result.Path = s.path
		result.Version = moduleVersion
		result.BootTime = &bootTime
		result.Remark = moduleRemark
		result.Interactive = service.Interactive()
	} else if argument.Action == "restart" {
		if len(strings.TrimSpace(argument.Name)) < 1 {
			result.Code = 3
			result.Error = fmt.Sprintf("inpunt invalid: service name '%s' is empty", argument.Name)
			result.Output(w)
			return
		}
		mgr := &Service{}
		err = mgr.Restart(argument.Name)
		if err != nil {
			result.Code = 4
			result.Error = fmt.Sprintf("restart service error: %v", err)
		}
	} else if argument.Action == "update" {
		if len(strings.TrimSpace(argument.Name)) < 1 {
			result.Code = 3
			result.Error = fmt.Sprintf("inpunt invalid: service name '%s' is empty", argument.Name)
			result.Output(w)
			return
		}

		if len(strings.TrimSpace(argument.Path)) < 1 {
			result.Code = 3
			result.Error = fmt.Sprintf("inpunt invalid: service execute file path '%s' is empty", argument.Path)
			result.Output(w)
			return
		}

		if len(strings.TrimSpace(argument.UpdateFile)) < 1 {
			result.Code = 3
			result.Error = fmt.Sprintf("inpunt invalid: service new execute file path '%s' is empty", argument.UpdateFile)
			result.Output(w)
			return
		}
		_, err := os.Stat(argument.UpdateFile)
		if os.IsNotExist(err) {
			result.Code = 3
			result.Error = fmt.Sprintf("inpunt invalid: service new execute file '%s' not exist", argument.UpdateFile)
			result.Output(w)
			return
		}

		svcUpd := &types.SvcUpd{
			Name: argument.Name,
			Mgr:  &Service{},
		}
		err = svcUpd.Update(argument.Path, argument.UpdateFile, argument.UpdateFolder)
		if err != nil {
			result.Code = 5
			result.Error = fmt.Sprintf("update service error: %v", err)
		}
	} else {
		result.Code = 3
		result.Error = fmt.Sprintf("inpunt invalid: action '%s' not support", argument.Action)
	}

	result.Output(w)
}
