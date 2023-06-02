package loghelper

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var File *os.File

// CreateFile Метод создания логов
func CreateFile() {
	//Проверка создана ли папка для логов
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			fmt.Println(err.Error())
			WriteLogs(strings.ReplaceAll(err.Error(), "pq: ", ""))
		}
	}

	//Создание логов
	file, err := os.Create("./logs/" + time.Now().Format("2006.01.02(15h04m05s)") + ".txt")
	File = file
	if err != nil {
		fmt.Println(err.Error())
		WriteLogs(strings.ReplaceAll(err.Error(), "pq: ", ""))
	}
}

// WriteLogs Запись строки в логи
func WriteLogs(text string) {
	_, err := File.WriteString(time.Now().Format("2006.01.02(15h04m05s)") + ": " + text + "\n")
	if err != nil {
		fmt.Println(err.Error())
		WriteLogs(strings.ReplaceAll(err.Error(), "pq: ", ""))
	}
	//defer File.Close()
}

// PrintHostParameters Запись в логи и вывод в терминале информации о параметрах хостов
func PrintHostParameters(hostid string, value string) {
	fmt.Printf("HostId: %v, Value: %v, Time_load: %v\n", hostid, value, time.Now().Format("2006.01.02(15h04m05s)"))
	WriteLogs(fmt.Sprintf("HostId: %v, Value: %v, Time_load: %v", hostid, value, time.Now().Format("2006.01.02(15h04m05s)")))
}
