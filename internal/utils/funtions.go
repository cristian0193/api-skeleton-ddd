package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(name string) (string, error) {
	s, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("env var %s not found", name)
	}
	return s, nil
}

func GetInt(name string) (int, error) {
	s, ok := os.LookupEnv(name)
	if !ok {
		return 0, fmt.Errorf("env var %s not found", name)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("env var %s must be a number", name)
	}
	return n, nil
}
