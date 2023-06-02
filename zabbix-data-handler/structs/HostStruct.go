package structs

type HostInterfaceStruct struct {
	IP string `json:"ip"`
}

// HostStruct Структура описания хоста
type HostStruct struct {
	Hostid   string                `json:"hostid"`
	HostName string                `json:"host"`
	HostIP   []HostInterfaceStruct `json:"interfaces"`
}

// HostsStruct Структура хоста
type HostsStruct struct {
	Jsonrpc string       `json:"jsonrpc"`
	Result  []HostStruct `json:"result"`
	Id      int          `json:"id"`
}

var Hosts HostsStruct
