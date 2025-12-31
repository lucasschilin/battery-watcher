package main

import (
	"os"
	"strings"
)

func isChargerConnected() (bool, error) {
	data, err := os.ReadFile(chargerPath)
	if err != nil {
		return false, err
	}

	value := strings.TrimSpace(string(data))

	return value == "1", nil
}
