package gu

import "log"

func LogFatalIfNotNil(err error, message string) {
	if err != nil {
		if message != "" {
			log.Fatal(message)
		} else {
			log.Fatal(err)
		}
	}
}

