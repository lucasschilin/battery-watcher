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

		fmt.Printf("current level = %d%%\n", level)
		isCharging, err := isChargerConnected()
		if err != nil {
			fmt.Println("Erro ao verificar carregador:", err)
			continue
		}
		fmt.Printf("charger connected = %v\n", isCharging)

		if level > lowLevelLimit && level < highLevelLimit {
			notifiedInLowLevelLimit = false
			notifiedInHighLevelLimit = false
		}

		if isCharging && level >= highLevelLimit && !notifiedInHighLevelLimit {
			err := sendNotification(
				"Bateria carregada",
				fmt.Sprintf("Bateria em %d%%. Considere desconectar o carregador.", level),
			)
			if err != nil {
				fmt.Println("Erro ao enviar notificação:", err)
				continue
			}

			notifiedInHighLevelLimit = true
		}

		if !isCharging && level <= lowLevelLimit && !notifiedInLowLevelLimit {
			err := sendNotification(
				"Bateria baixa",
				fmt.Sprintf("Bateria em %d%%. Considere conectar o carregador.", level),
			)
			if err != nil {
				fmt.Println("Erro ao enviar notificação:", err)
				continue
			}

			notifiedInLowLevelLimit = true
		}

		time.Sleep(sleepTimeInSeconds * time.Second)
	}
}
