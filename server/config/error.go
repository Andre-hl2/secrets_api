package config

import "fmt"

type MissingEnvFile struct{}

func (me MissingEnvFile) Error() string {
	return "Error loading .env file"
}

type MissingEnvVar struct {
	EnvVar string
}

func (me MissingEnvVar) Error() string {
	return fmt.Sprintf("Missing env var '%s'", me.EnvVar)
}

type IncorrectEnvVarValue struct {
	EnvVar  string
	Options []string
}

func (ie IncorrectEnvVarValue) Error() string {
	return fmt.Sprintf("Invalid value for '%s' var. Should be one of: %s", ie.EnvVar, ie.Options)
}
