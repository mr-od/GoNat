package main

import (
	"flag"

	// "github.com/mr-od/a4chat/util/validation"

	// "github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {

	flag.Parse()

	// binding.Validator = validation.NewDefaultValidator()

}
