package environment

import "os"

// Application Environment name
const (
	Development = "development"
	Test        = "test"
	E2E         = "e2e"
)

// IsDev returns APP_ENV in development mode
func IsDev() bool {
	println("os.Getenv(APP_ENV): ", os.Getenv("APP_ENV"))
	return os.Getenv("APP_ENV") == Development
}

// IsTest returns APP_ENV in test mode
func IsTest() bool {
	return os.Getenv("APP_ENV") == Test
}

// IsE2E returns APP_ENV in e2e mode
func IsE2E() bool {
	return os.Getenv("APP_ENV") == E2E
}
