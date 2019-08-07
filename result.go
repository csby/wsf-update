package main

import (
	"encoding/json"
	"fmt"
	"github.com/csby/wsf/types"
	"net/http"
)

type Result struct {
	types.SvcUpdResult
}

func (s *Result) Output(w http.ResponseWriter) {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(data)
	}
}
