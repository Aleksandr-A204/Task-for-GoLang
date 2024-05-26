package models

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
