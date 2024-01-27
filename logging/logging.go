package logging

import (
	"log"
	"os"
)

func LogUserActivity(userID, activity string) {
	logFile, err := os.OpenFile("logging/user_activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Printf("The user with the ID %s %s\n", userID, activity)
}

func LogUserAuth(activity string) {
	logFile, err := os.OpenFile("logging/user_activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	logger.Printf("The user is %s\n", activity)
}
