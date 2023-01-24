package ROCKSError

import (
	"log"
)

type ErrorInformation struct {
	ErrorType string
	Error     error
}

func Encounter(err ErrorInformation) {
	if err.Error != nil {
		switch err.ErrorType {
		case "FATAL":
			log.Fatalf("ROCKS.FATAL: %v", err.Error.Error())
		case "Warning":
			log.Printf("ROCKS.WARNING: %v", err.Error.Error())
		case "INFO":
			log.Printf("ROCKS.INFO: %v", err.Error.Error())
		default:
			return
		}
	}
	return
}
