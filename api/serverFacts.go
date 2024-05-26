package api

import (
	"bytes"
	"encoding/json"
	"task/cmd/models"
)

// Функция для сохранения записей из буфера в БД, которая сохраняет данные в БД
func SaveToBD(factData models.FactData) error {
	resource := "/facts/save_fact" // определение ресурс API для сохранения данных

	// Сериализацуем данные факта в JSON
	jsonBody, err := json.Marshal(factData)
	if err != nil {
		return err
	}

	// Создаем новый HTTP-запрос с сериализованным телом JSON
	err = serverHttp(resource, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return nil
}
