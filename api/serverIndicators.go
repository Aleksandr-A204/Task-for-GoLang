package api

import (
	"bytes"
	"encoding/json"
	"task/cmd/models"
)

// Функция для сохранения записей из буфера в БД, которая сохраняет данные в БД
func GetFacts() error {
	resource := "/indicators/get_facts" // определение ресурс API для сохранения данных

	factData := models.FactsIndicators{
		PeriodStart:     "2024-05-01",
		PeriodEnd:       "2023-05-31",
		PeriodKey:       "month",
		IndicatorToMoID: "227373",
	}

	// Сериализацуем данные факта в JSON
	jsonBody, err := json.Marshal(factData)
	if err != nil {
		return err
	}

	err = serverHttp(resource, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return nil
}
