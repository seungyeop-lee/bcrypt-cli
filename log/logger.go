package logger

import "log"

var IsVerbose = false

func Info(message string) {
	if IsVerbose {
		log.Println(message)
	}
}

func Error(error error) {
	log.Println(error)
}
