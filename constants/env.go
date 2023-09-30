package constants

import "os"

var (
	REDIS_HOST     = os.Getenv("REDIS_HOST")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	POSTGRE        = os.Getenv("POSTGRE")
)
