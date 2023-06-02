package configHelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"zabbix-data-handler/errorHelper"
	"zabbix-data-handler/structs"
)

// GetConfig Получение конфига
func GetConfig() {
	var err error
	if _, err = os.Stat("./cfg.json"); os.IsNotExist(err) {
		//Создание файла конфига
		fmt.Printf("Создание файла конфигурации...")

		configInfo, err := json.MarshalIndent(
			structs.ConfigStruct{
				LoginZabbix:    "Admin",
				PasswordZabbix: "zabbix",
				Time:           "00h01m00s",
				Url:            "http://192.168.42.200/zabbix/api_jsonrpc.php",
				HostDB:         "80.78.241.80",
				PortDB:         5433,
				UserDB:         "postgres",
				PasswordDB:     "password",
				DBName:         "zabbix_api",
				LoginServer:    "user",
				PasswordServer: "12345678",
				DoDeleteFrom:   true,
			}, "", "\t")
		errorHelper.PrintError(err)
		err = ioutil.WriteFile("./cfg.json", configInfo, 0644)
		errorHelper.PrintError(err)
	}
	//Загрузка файлов конфигурации
	fmt.Printf("Загрузка файла конфигурации...")

	configInfo, err := ioutil.ReadFile("./cfg.json")
	errorHelper.PrintError(err)
	err = json.Unmarshal(configInfo, &structs.Config)
	errorHelper.PrintError(err)
}
