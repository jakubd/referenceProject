package main

import (
	"fmt"
	ref "github.com/jakubd/referenceProject"
	log "github.com/sirupsen/logrus"
)

func setUpLogger() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.SetLevel(log.InfoLevel)
}

func main() {

	setUpLogger()
	log.Info("started main func")

	a := 1
	b := 2
	fmt.Printf("%d + %d = %d\n", a, b, ref.AddNumbers(a, b))

	log.Info("ended main func")
}
