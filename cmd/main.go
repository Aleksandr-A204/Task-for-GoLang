package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Структура для хранения данных факта
type FactData struct {
	PeriodStart         string `json:"period_start"`
	PeriodEnd           string `json:"period_end"`
	PeriodKey           string `json:"period_key"`
	IndicatorToMoID     string `json:"indicator_to_mo_id"`
	IndicatorToMoFactID string `json:"indicator_to_mo_fact_id"`
	Value               string `json:"value"`
	FactTime            string `json:"fact_time"`
	IsPlan              string `json:"is_plan"`
	AuthUserID          string `json:"auth_user_id"`
	Comment             string `json:"comment"`
}

var url string = "https://development.kpi-drive.ru/_api/" // URL API для сохранения данных факта
var token string = "48ab34464a5573519725deb5865cc74c"     // Токен для авторизации в API

func main() {
	buffer := make(chan FactData, 1000) // Инициализируем буфер емкостью 1000
	then := time.Now()

	// Создаем контекст для управления жизненным циклом горутин
	ctx := context.Background()
	// Запустим горутину для чтения данных из буффера и их сохранения в БД
	go func() {
		for {
			select {
			case <-ctx.Done(): // Проверка отмены контекста
				return // Выход из горутины
			case factData := <-buffer: // Получаем данные из буфера
				fmt.Println(factData)
				// Попытаемся сохранить данные в БД и логирование ошибки, если таковая возникает.
				if err := saveToBD(factData); err != nil {
					log.Println(err)
				}
			}
		}
	}()

	// Инициализируем группу ожидания для синхронизации завершения горутин
	var wg sync.WaitGroup
	// Установим счётчик группы ожидания в 10, что соответствует количеству горутин
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done() // Уменьшение счётчика группы ожидания после завершения горутины.
			factData := FactData{
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
	// Закроем буффер по завершению сохранения данных
	close(buffer)

	elapsed := time.Since(then)
	log.Println("Total time taken:", elapsed)
}

// Функция для сохранения записей из буфера в БД, которая сохраняет данные в БД
func saveToBD(fact FactData) error {
	resource := "facts/save_fact" // определение ресурс API для сохранения данных

	// Сериализацуем данные факта в JSON
	jsonBody, err := json.Marshal(fact)
	if err != nil {
		return err
	}

	// Создаем новый HTTP-запрос с сериализованным телом JSON
	req, err := http.NewRequest("POST", url+resource, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	// Установим заголовки запроса: тип содержимого и авторизационный токен.
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Создаем HTTP-клиент и выполним запрос.
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}
