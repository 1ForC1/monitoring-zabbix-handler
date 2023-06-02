package DB

import (
	"database/sql"
	"fmt"
	"zabbix-data-handler/errorHelper"
	_ "zabbix-data-handler/github.com/lib/pq"
	"zabbix-data-handler/logHelper"
	"zabbix-data-handler/structs"
)

var db *sql.DB
var err error

// DbConnect Метод подключения к БД
func DbConnect() {
	//Строка подключения
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", structs.Config.HostDB, structs.Config.PortDB, structs.Config.UserDB, structs.Config.PasswordDB, structs.Config.DBName)
	//Подключение к БД
	db, err = sql.Open("postgres", psqlconn)
	errorHelper.PrintError(err)
	//defer db.Close()

	//Пинг БД
	err = db.Ping()
	if err != nil {
		errorHelper.PrintError(err)
	} else {
		fmt.Println("\nУспешное подключение!")
		loghelper.WriteLogs("Успешное подключение!")
	}
}
