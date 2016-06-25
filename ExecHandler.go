package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func (m triggerElement) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	splitted := strings.Split(r.RequestURI, "/")
	if len(splitted) == 3 {
		if splitted[2] == "trigger" {
			log.Printf("Route %s called, executing %s", m.Route, m.Executable)
			if m.CheckExecutable {
				if _, err := os.Stat(m.Executable); os.IsNotExist(err) {
					log.Println("Executable not found")
					if j, err := json.Marshal(triggerResult{true, "Executable not found"}); err == nil {
						w.Write(j)
					}
					return
				}
			}
			go execDeffered(m.Executable, m)
			if j, err := json.Marshal(triggerResult{true, ""}); err == nil {
				w.Write(j)
			}
		} else {
			if val, ok := logs[m.Route]; ok {
				if j, err := json.Marshal(logResult{true, val}); err == nil {
					w.Write(j)
				}
			} else {
				w.WriteHeader(404)
			}
		}
	} else {
		w.WriteHeader(500)
	}
}

func execDeffered(executable string, route triggerElement) {
	splitted := strings.SplitN(executable, " ", 2)
	cmd := exec.Command(splitted[0], splitted[1])
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	log.Printf("Command finished with error: %v", err)
	logs[route.Route] = string(out)
}
