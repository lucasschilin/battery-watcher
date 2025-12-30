package main

import "os/exec"

func sendNotification(title, message string) error {
	cmd := exec.Command("notify-send", title, message)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
