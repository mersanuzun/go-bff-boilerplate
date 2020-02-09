package providers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Environment() Provider {
	return &environmentConfig{}
}

type environmentConfig struct{}

func (c *environmentConfig) GetString(key string) (string, error) {
	return getEnvironmentVariable(key)
}

func (c *environmentConfig) GetInt(key string) (int, error) {
	valueStr, err := getEnvironmentVariable(key)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(valueStr)
}

func (c *environmentConfig) GetBool(key string) (bool, error) {
	valueStr, err := getEnvironmentVariable(key)

	if err != nil {
		return false, err
	}

	return strconv.ParseBool(valueStr)
}

func (c *environmentConfig) GetMap(key string) (map[string]string, error) {
	return nil, errors.New("Not supported") // TODO add support for Map parameters via flattening in the future
}

func getEnvironmentVariable(variableName string) (string, error) {
	environmentVariable := os.Getenv(variableName)

	if environmentVariable == "" {
		return "", errors.New(fmt.Sprintf("Environment Variable (%s) is not provided", variableName))
	}

	return environmentVariable, nil
}
