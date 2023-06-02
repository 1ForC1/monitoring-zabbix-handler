package apiRequests

import (
	"encoding/json"
	"zabbix-data-handler/errorHelper"
	"zabbix-data-handler/structs"
)

// GetTriggers Метод обработки полученных данных из zabbix api по загрузке триггеров
func GetTriggers(hostid string) bool {
	//Запрос
	var jsonRequest = []byte(`{
    "jsonrpc": "2.0",
    "method": "trigger.get",
    "params": {
        "output": [
			"triggerid",
            "expression",
            "description",
            "priority"
        ],
        "filter": {
            "value": "1",
            "hostid": ` + hostid + `
        },
		"sortfield": "priority",
		"sortorder": "DESC",
        "expandExpression": true,
        "expandDescription": true,
        "expandComment": true
    },
    "id": 1,
    "auth": "` + structs.Auth.Token + `"
}`)

	//Получение тела из шаблона обработки запросов
	body := GetInfoTemplate("GET", "", jsonRequest)
	//Запись тела в структуру
	if body != nil {
		err := json.Unmarshal(body, &structs.Triggers)
		if err != nil {
			errorHelper.PrintError(err)
			return false
		}
		return true
	}
	return false
}
