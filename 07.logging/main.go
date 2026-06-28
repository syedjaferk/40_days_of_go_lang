package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Config Values")
	fmt.Println(cfg)

	log.WithField("app", cfg.App.Name).Info("Application Ended")
	log.WithFields(log.Fields{
		"user": "Syed Jafer",
		"age":  28,
	}).Warn("User Validation Failed")
}
