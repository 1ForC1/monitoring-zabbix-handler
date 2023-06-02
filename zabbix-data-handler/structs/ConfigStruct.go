package structs

// ConfigStruct Структура конфигурации
type ConfigStruct struct {
	LoginZabbix    string
	PasswordZabbix string
	Time           string
	Url            string
	HostDB         string
	PortDB         int
	UserDB         string
	PasswordDB     string
	DBName         string
	LoginServer    string
	PasswordServer string
	DoDeleteFrom   bool
}

// Config Данные по умолчанию
var Config = ConfigStruct{}
