package main

type triggerElement struct {
	Route           string
	Executable      string
	CheckExecutable bool
}

type configDocument struct {
	Adress  string
	Trigger []triggerElement
}

type triggerResult struct {
	Status bool `json:"status"`
}

type logResult struct {
	Status bool   `json:"status"`
	Log    string `json:"log"`
}
