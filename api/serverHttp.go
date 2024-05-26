package api

import (
	"fmt"
	"io"
	"net/http"
)

func serverHttp(resource string, body io.Reader) error {
	// Создаем новый HTTP-запрос с сериализованным телом JSON
	request, err := http.NewRequest("POST", "https://development.kpi-drive.ru/_api"+resource, body)
	if err != nil {
		return err
	}

	// Установим заголовки запроса: тип содержимого и авторизационный токен.
	request.Header.Set("Authorization", "Bearer 48ab34464a5573519725deb5865cc74c")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Создаем HTTP-клиент и выполним запрос.
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	return nil
}
