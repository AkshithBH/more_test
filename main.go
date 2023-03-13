package more_test

import "log"

func logInfo(message string) {
	log.Println("info - %v", message)
}
func logWarning(message string) {
	log.Println("warn - %v", message)
}
func logError(message string) {
	log.Println("error - %v", message)
}
