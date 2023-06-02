package structs

// ParameterResult Структура информации о параметрах хостов
type ParameterResult struct {
	Itemid string `json:"itemid"`
	HostId string `json:"hostid"`
	Value  string `json:"lastvalue"`
}

// ParameterStruct Структура параметров хостов
type ParameterStruct struct {
	Jsonrpc string            `json:"jsonrpc"`
	Result  []ParameterResult `json:"result"`
	Id      int               `json:"id"`
}

var CpuLoad1 ParameterStruct
var CpuLoad5 ParameterStruct
var CpuLoad15 ParameterStruct
var RamFree ParameterStruct
var RamTotal ParameterStruct
var DiskSpaceFree ParameterStruct
var DiskSpaceTotal ParameterStruct
var CpuUtilUser ParameterStruct
var CpuUtilIdle ParameterStruct
var CpuUtilSystem ParameterStruct
var CpuUtilIowait ParameterStruct
var CpuUtilNice ParameterStruct
var CpuUtilInterrupt ParameterStruct
var CpuUtilSoftirq ParameterStruct
var CpuUtilSteal ParameterStruct
