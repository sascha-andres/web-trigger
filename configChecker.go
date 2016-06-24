package main

import (
	"errors"
)

func (cfg configDocument) checkConfig() (map[string]triggerElement, error) {
	triggers := make(map[string]triggerElement)
	for _, item := range cfg.Trigger {
		if _, ok := triggers[item.Route]; ok {
			return triggers, errors.New("Route provided multiple times")
		}
		triggers[item.Route] = item
	}
	return triggers, nil
}
