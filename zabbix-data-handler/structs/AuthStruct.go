package structs

// AuthStruct Структура авторизации в zabbix
type AuthStruct struct {
	Jsonrpc string `json:"jsonrpc"`
	Token   string `json:"result"`
	Id      int    `json:"id"`
}

var Auth AuthStruct
