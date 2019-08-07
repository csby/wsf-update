package main

type Argument struct {
	Action       string `json:"action" note:"info or update"`
	Name         string `json:"name" note:"service name"`
	Path         string `json:"path" note:"execute file path"`
	UpdateFolder string `json:"updateFolder" note:"temp folder to be deleted"`
	UpdateFile   string `json:"updateFile" note:"path of the new execute file"`
}
