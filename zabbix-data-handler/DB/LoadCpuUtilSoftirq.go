package DB

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"zabbix-data-handler/apiRequests"
	"zabbix-data-handler/errorHelper"
	loghelper "zabbix-data-handler/logHelper"
	"zabbix-data-handler/structs"
)

func LoadCpuUtilSoftirq() {
	fmt.Printf("Загрузка данных заргуженности процессора(прогаммные прерывания в процентах)...\n")
	loghelper.WriteLogs("Загрузка данных заргуженности процессора(прогаммные прерывания в процентах)...")
	body := apiRequests.GetInfoTemplate("GET", "system.cpu.util[,softirq]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.CpuUtilSoftirq)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.CpuUtilSoftirq.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.CpuUtilSoftirq.Result[i].HostId {
					loghelper.PrintHostParameters(structs.CpuUtilSoftirq.Result[i].HostId, structs.CpuUtilSoftirq.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.CpuUtilSoftirq.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					value, err := strconv.ParseFloat(structs.CpuUtilSoftirq.Result[i].Value, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "cpu_softirq" ("hostid", "value_cpu_softirq", "time_cpu_softirq") values (%d , %v , '%v');`+"\n", hostid, errorHelper.Round(value, 0.0005), time.Now().Format("2006-01-02 15:04:05"))
					errorHelper.PrintError(err)
					break
				}
			}
		}
	} else {
		fmt.Printf("Пустой ответ")
		loghelper.WriteLogs(fmt.Sprintf("Пустой ответ"))
	}
}
