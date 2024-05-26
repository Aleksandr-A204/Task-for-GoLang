package components

import (
	"context"
	"fmt"
	"log"
	"task/api"
	"task/cmd/models"
)

// Главная функция для обработки буфера и сохранения записей
func ProcessBuffer(buffer chan models.FactData) {
	// Создаем контекст для управления жизненным циклом горутин
	ctx := context.Background()

	for {
		select {
		case <-ctx.Done(): // Проверка отмены контекста
			return // Выход из горутины
		case factData := <-buffer: // Получаем данные из буфера
			fmt.Println(factData)
			// Попытаемся сохранить данные в БД и логирование ошибки, если таковая возникает.
			if err := api.SaveToBD(factData); err != nil {
				log.Println(err)
			}
		}
	}
}
