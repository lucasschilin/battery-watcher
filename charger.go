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

	// Remove espaços e quebra de linha
	value := strings.TrimSpace(string(data))

	// "1" = conectado | "0" = desconectado
	return value == "1", nil
}
