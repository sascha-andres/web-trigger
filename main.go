package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var hostName string
var triggers map[string]triggerElement
var logs map[string]string

func main() {
	var configPath string

	args := os.Args[1:]
	numberOfArgs := len(args)
	hostName, _ = os.Hostname()
	var config configDocument
	logs = make(map[string]string)

	log.Println("Configuration ...")
	if 0 == numberOfArgs {
		if _, err := readConfig("web-trigger.yml"); nil == err {
			configPath = "web-trigger.yml"
		}
	} else {
		if 1 != numberOfArgs {
			log.Fatal("Please provide config file")
			os.Exit(1)
		} else {
			configPath = args[0]
		}
	}
	log.Printf("Using '%s' as config", configPath)

	if cfg, err := readConfig(configPath); nil != err {
		log.Fatal(err)
		os.Exit(2)
	} else {
		config = cfg
	}

	if trg, err := config.checkConfig(); nil != err {
		log.Fatal(err)
		os.Exit(3)
	} else {
		triggers = trg
	}

	log.Println("Setting up routes ...")

	r := mux.NewRouter().StrictSlash(false)

	for _, route := range triggers {
		r.Handle(fmt.Sprintf("/%s/trigger", route.Route), loggingHandler(route)).Methods("GET")
		r.Handle(fmt.Sprintf("/%s/log", route.Route), loggingHandler(route)).Methods("GET")
	}

	log.Printf("Listening on %s ...", config.Adress)
	server := &http.Server{
		Addr:    config.Adress,
		Handler: r,
	}
	server.ListenAndServe()
}
