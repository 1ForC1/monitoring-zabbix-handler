package apiRequests

import (
	"encoding/json"
	"zabbix-data-handler/errorHelper"
	"zabbix-data-handler/structs"
)

// PostAuth Метод отправки данных на zabbix для авторизации
func PostAuth() bool {
	//Запрос
	var jsonRequest = []byte(`{
    "jsonrpc": "2.0",
    "method": "user.login",
    "params": {
        "user": "` + structs.Config.LoginZabbix + `",
		"password": "` + structs.Config.PasswordZabbix + `"
    },
    "id": 1
}`)

	//Получение тела из шаблона обработки запросов
	body := GetInfoTemplate("POST", "", jsonRequest)
	//Запись тела в структуру
	err := json.Unmarshal(body, &structs.Auth)
	if err != nil {
		errorHelper.PrintError(err)
		return false
	}
	return true
}
