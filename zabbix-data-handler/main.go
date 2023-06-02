package main

import (
	"fmt"
	"time"
	"zabbix-data-handler/DB"
	"zabbix-data-handler/apiRequests"
	"zabbix-data-handler/configHelper"
	"zabbix-data-handler/errorHelper"
	loghelper "zabbix-data-handler/logHelper"
	"zabbix-data-handler/structs"
)

// Главный метод
func main() {
	//Создание логов
	loghelper.CreateFile()
	//Получение или создание конфигурации
	configHelper.GetConfig()
	//Расчёт частоты обновления данных
	timer, err := time.ParseDuration(structs.Config.Time)
	errorHelper.PrintError(err)

	//Подключение к БД
	fmt.Printf("\nПодключение к базе данных...")
	loghelper.WriteLogs("Подключение к базе данных...")
	DB.DbConnect()

	//Авторизация в zabbix
	fmt.Printf("\nАвторизация...")
	loghelper.WriteLogs("Авторизация...")
	if apiRequests.PostAuth() {
		for {
			//Обновление таблиц в БД
			fmt.Println("\nОбновление данных...")
			loghelper.WriteLogs("Обновление данных...")
			err = DB.ReloadTables()
			if err != nil {
				fmt.Printf("Данные не обновлены\n\n")
				loghelper.WriteLogs("Данные не обновлены\n\n")
			} else {
				fmt.Printf("Данные обновлены\n\n")
				loghelper.WriteLogs("Данные обновлены\n\n")
			}

			//Таймер на переотправку запросов
			time.Sleep(timer)
		}
	} else {
		fmt.Printf("Произошла ошибка при подключении к серверу zabbix...")
		loghelper.WriteLogs("Произошла ошибка при подключении к серверу zabbix.")
	}
}
