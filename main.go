package main

import (
	"fmt"
)

func main() {
	var status string
	fmt.Println("Сканер логов запущен. Вводите статусы (ping, error, exit или любые слова):")

MyLoop:
	for {
		fmt.Scan(&status)

		switch status {
		case "ping":
			continue MyLoop
		case "error":
			fmt.Println("Обнаружена ошибка! Проверяем систему...")
		case "exit":
			fmt.Println("Выход из сканера.")
			break MyLoop
		default:
			fmt.Printf("Лог записан: %s.\n", status)
		}
	}
}
