package main

import (
	"fmt"
	"log"
	"task/cmd/components"
	"task/cmd/models"
	"time"
)

func main() {
	fmt.Println("Application is running...")
	fmt.Println("Press any key to end the application.")
	fmt.Println()

	buffer := make(chan models.FactData, 1000) // Инициализируем буфер емкостью 1000
	then := time.Now()

	// Запустим горутину для чтения данных из буффера и их сохранения в БД
	go components.ProcessBuffer(buffer)

	// Добавим входящие данные факта в буффер
	components.AddToBuffer(buffer)

	//close(buffer) // Закроем буффер по завершению сохранения данных

	elapsed := time.Since(then)
	log.Println("Total time taken:", elapsed)
	fmt.Scanln()
}
