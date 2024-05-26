package components

import (
	"fmt"
	"sync"
	"task/cmd/models"
)

// Функция обработки входящих данных и добавления их в буфер
func AddToBuffer(buffer chan models.FactData) {
	// Инициализируем группу ожидания для синхронизации завершения горутин
	var wg sync.WaitGroup
	// Установим счётчик группы ожидания в 10, что соответствует количеству горутин
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done() // Уменьшение счётчика группы ожидания после завершения горутины.
			factData := models.FactData{
				PeriodStart:         fmt.Sprintf("2024-05-%02d", i+1),
				PeriodEnd:           "2024-05-31",
				PeriodKey:           "month",
				IndicatorToMoID:     "227373",
				IndicatorToMoFactID: "0",
				Value:               "1",
				FactTime:            fmt.Sprintf("2024-05-%02d", i+1),
				IsPlan:              "0",
				AuthUserID:          "40",
				Comment:             fmt.Sprintf("buffer Last_name-%d", i),
			}
			buffer <- factData // Сохраним сгенерированные данные факта в буффер
		}(i)
	}

	wg.Wait() // Ожидаем завершения всех сохраненных горутин
}
