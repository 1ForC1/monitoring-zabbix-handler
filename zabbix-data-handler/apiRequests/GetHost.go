package apiRequests

import (
	"encoding/json"
	"zabbix-data-handler/errorHelper"
	"zabbix-data-handler/structs"
)

// GetHost Метод обработки полученных данных из zabbix api по загрузке хостов
func GetHost() {
	//Запрос
	var jsonRequest = []byte(`{
    "jsonrpc": "2.0",
    "method": "host.get",
    "params": {
        "output": [
            "host"
        ],
        "selectInterfaces": [
            "ip"
        ],
		"sortfield": "hostid",
        "sortorder": "ASC"
    },
    "id": 1,
    "auth": "` + structs.Auth.Token + `"
}`)

	//Получение тела из шаблона обработки запросов
	body := GetInfoTemplate("GET", "", jsonRequest)
	//Запись тела в структуру
	err := json.Unmarshal(body, &structs.Hosts)
	errorHelper.PrintError(err)
}
