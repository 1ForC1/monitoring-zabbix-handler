package structs

// TriggerStruct Структура описания триггера
type TriggerStruct struct {
	Triggerid   string `json:"triggerid"`
	Expression  string `json:"expression"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

// TriggersStruct Структура тиггера
type TriggersStruct struct {
	Jsonrpc string          `json:"jsonrpc"`
	Result  []TriggerStruct `json:"result"`
	Id      int             `json:"id"`
}

var Triggers TriggersStruct
