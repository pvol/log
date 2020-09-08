package main

import "github.com/pvol/log"

func execColorizedExamples() {
	log.SetLevel(log.Lall)
	log.Info("default config")

	log.Colorized(true)
	log.Info("colorized config")

	log.Colorized(false)
	log.Error("close colorized config")
}
