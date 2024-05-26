package models

// Структура для проверки метод получения записей с сервера
type FactsIndicators struct {
	PeriodStart     string `json:"period_start"`
	PeriodEnd       string `json:"period_end"`
	PeriodKey       string `json:"period_key"`
	IndicatorToMoID string `json:"indicator_to_mo_id"`
}
