package helper

import "log"

func Error(err error) {
	if err != nil {
		log.Fatal("ERROR CONNECT")
	}

}
