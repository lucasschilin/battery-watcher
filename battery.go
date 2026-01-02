package main

import (
	"os"
	"strconv"
	"strings"
)

func readBatteryLevel(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	value := strings.TrimSpace(string(data))

	level, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return level, nil
}
