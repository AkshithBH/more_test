package more_test

import "log"

func LogInfo(message string) {
	log.Println("info - %v", message)
}
func LogWarning(message string) {
	log.Println("warn - %v", message)
}
func LogError(message string) {
	log.Println("error - %v", message)
}
