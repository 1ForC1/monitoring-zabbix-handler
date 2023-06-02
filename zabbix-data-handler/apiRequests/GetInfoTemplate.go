package apiRequests

import (
	"bytes"
	"fmt"
	"net/http"
	"zabbix-data-handler/errorHelper"
	"zabbix-data-handler/structs"
)

// GetInfoTemplate Шаблон отправки и получения zabbix api
func GetInfoTemplate(requestType string, key string, jsonRequest []byte) []byte {
	if jsonRequest == nil {
		jsonRequest = []byte(`{
    "jsonrpc": "2.0",
    "method": "item.get",
    "params": {
		"output": [
			"hostid",
            "lastvalue"
        ],
        "search": {
            "key_": "` + key + `"
        },
        "softfield": "name"
    },
    "id": 1,
    "auth": "` + structs.Auth.Token + `"
}`)
	}

	//Создание нового запроса
	req, err := http.NewRequest(requestType, structs.Config.Url, bytes.NewBuffer(jsonRequest))
	errorHelper.PrintError(err)
	req.Header.Set("Content-Type", "application/json-rpc")

	client := &http.Client{}
	//Базовая авторизация для сервера
	req.SetBasicAuth("user", "12345678")

	//Отправка запроса
	resp, err := client.Do(req)
	if err != nil {
		errorHelper.PrintError(err)
		return nil
	}
	defer resp.Body.Close()
	fmt.Println("Статус запроса: ", resp.Status)
	//Возвращается на обработку ошибок в ResponseBody
	return ResponseBody(resp)
}
