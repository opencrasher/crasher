package app

import "os"

var (
	ServerAddress = os.Getenv("CRASHER_SERVER_ADDRESS")

	DatabaseAddress  = os.Getenv("CRASHER_DATABASE_ADDRESS")
	DatabaseUsername = os.Getenv("CRASHER_DATABASE_USERNAME")
	DatabasePassword = os.Getenv("CRASHER_DATABASE_PASSWORD")
)
