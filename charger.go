package main

import (
	"os"
	"strings"
)

func isChargerConnected(path string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	value := strings.TrimSpace(string(data))

	return value == "1", nil
}
