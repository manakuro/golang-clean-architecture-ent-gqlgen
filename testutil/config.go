package testutil

import (
	"golang-clean-architecture-ent-gqlgen/config"
	"golang-clean-architecture-ent-gqlgen/pkg/util/environment"
)

// ReadConfig reads config file for test.
func ReadConfig() {
	config.ReadConfig(config.ReadConfigOption{
		AppEnv: environment.Test,
	})
}

// ReadConfigE2E reads config file for e2e.
func ReadConfigE2E() {
	config.ReadConfig(config.ReadConfigOption{
		AppEnv: environment.E2E,
	})
}
