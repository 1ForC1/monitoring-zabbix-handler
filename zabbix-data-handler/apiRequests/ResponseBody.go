package apiRequests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"zabbix-data-handler/errorHelper"
	"zabbix-data-handler/structs"
)

// ResponseBody Обработка полученного ответа от запроса в zabbix
func ResponseBody(resp *http.Response) []byte {
	body, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(body, &structs.Error)
	errorHelper.PrintError(err)
	if structs.Error.Error.Code != 0 {
		//Вывод тела ошибки запроса
		fmt.Printf("Код: %v, Сообщение: %v, Описание: %v\n", structs.Error.Error.Code, structs.Error.Error.Message, structs.Error.Error.Data)
		structs.Error = structs.ErrorsStruct{}
		return nil
	}
	if strings.Contains(string(body), "\"result\":[]") {
		return nil
	} else {
		//Успешное возвращение тела ответа
		return body
	}
}
