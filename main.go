package main

import (
	"fmt"
	"time"
)

const (
	batteryPath        = "/sys/class/power_supply/BAT0/capacity"
	chargerPath        = "/sys/class/power_supply/AC/online"
	sleepTimeInSeconds = 60
	highLevelLimit     = 80
	lowLevelLimit      = 30
)

var (
	notifiedInLowLevelLimit  bool
	notifiedInHighLevelLimit bool
	lastSavedBatteryLevel    int
)

func main() {
	fmt.Println("🔋 Battery watcher iniciado")

	for {
		fmt.Println("\n=== NOVA LEITURA ===")
		level, err := readBatteryLevel()
		if err != nil {
			fmt.Println("Erro ao ler bateria:", err)
			time.Sleep(30 * time.Second)
			continue
		}
		if lastSavedBatteryLevel != 0 {
			fmt.Printf("last level = %d%%\n", lastSavedBatteryLevel)
		} else {
			fmt.Printf("last level = --\n")
		}
		fmt.Printf("current level = %d%%\n", level)
		isCharging, err := isChargerConnected()
		if err != nil {
			fmt.Println("Erro ao verificar carregador:", err)
			continue
		}
		fmt.Printf("charger connected = %v\n", isCharging)

		if isCharging && level >= highLevelLimit && !notifiedInHighLevelLimit && lastSavedBatteryLevel != 0 && level > lastSavedBatteryLevel {
			err := sendNotification(
				"Bateria carregada",
				fmt.Sprintf("Bateria em %d%%. Considere desconectar o carregador.", level),
			)
			if err != nil {
				fmt.Println("Erro ao enviar notificação:", err)
			}

			notifiedInHighLevelLimit = true
			notifiedInLowLevelLimit = false
		}

		if !isCharging && level <= lowLevelLimit && !notifiedInLowLevelLimit && lastSavedBatteryLevel != 0 && level < lastSavedBatteryLevel {
			err := sendNotification(
				"Bateria baixa",
				fmt.Sprintf("Bateria em %d%%. Considere conectar o carregador.", level),
			)
			if err != nil {
				fmt.Println("Erro ao enviar notificação:", err)
			}

			notifiedInLowLevelLimit = true
			notifiedInHighLevelLimit = false
		}

		lastSavedBatteryLevel = level
		time.Sleep(sleepTimeInSeconds * time.Second)
	}
}
