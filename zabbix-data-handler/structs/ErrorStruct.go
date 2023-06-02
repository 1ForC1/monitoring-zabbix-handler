package structs

// Description Структура описания ошибки
type Description struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// ErrorsStruct Структура ошибки
type ErrorsStruct struct {
	Jsonrpc string      `json:"jsonrpc"`
	Error   Description `json:"error"`
	Id      int         `json:"id"`
}

var Error ErrorsStruct
