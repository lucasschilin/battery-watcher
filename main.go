package main

import (
	"fmt"
	"time"

	"github.com/lucasschilin/battery-watcher/config"
)

var (
	notifiedInLowLevelLimit  bool
	notifiedInHighLevelLimit bool
)

func main() {
	fmt.Println("🔋 Battery watcher iniciado")

	for {
		config := config.LoadConfig()

		fmt.Println("\n=== NOVA LEITURA ===")
		level, err := readBatteryLevel(config.Battery.Path)
		if err != nil {
			fmt.Println("Erro ao ler bateria:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		fmt.Printf("current level = %d%%\n", level)
		isCharging, err := isChargerConnected(config.Charger.Path)
		if err != nil {
			fmt.Println("Erro ao verificar carregador:", err)
			continue
		}
		fmt.Printf("charger connected = %v\n", isCharging)

		if level > config.Battery.LowLevelLimit && level < config.Battery.HighLevelLimit {
			notifiedInLowLevelLimit = false
			notifiedInHighLevelLimit = false
		}

		if isCharging && level >= config.Battery.HighLevelLimit && !notifiedInHighLevelLimit {
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

		if !isCharging && level <= config.Battery.LowLevelLimit && !notifiedInLowLevelLimit {
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

		time.Sleep(time.Duration(config.SleepTimeInSeconds) * time.Second)
	}
}
