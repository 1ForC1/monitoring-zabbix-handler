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

func LoadCpuUtilInterrupt() {
	fmt.Printf("Загрузка данных заргуженности процессора(прерывания в процентах)...\n")
	loghelper.WriteLogs("Загрузка данных заргуженности процессора(прерывания в процентах)...")
	body := apiRequests.GetInfoTemplate("GET", "system.cpu.util[,interrupt]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.CpuUtilInterrupt)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.CpuUtilInterrupt.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.CpuUtilInterrupt.Result[i].HostId {
					loghelper.PrintHostParameters(structs.CpuUtilInterrupt.Result[i].HostId, structs.CpuUtilInterrupt.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.CpuUtilInterrupt.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					value, err := strconv.ParseFloat(structs.CpuUtilInterrupt.Result[i].Value, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "cpu_interrupt" ("hostid", "value_cpu_interrupt", "time_cpu_interrupt") values (%d , %v , '%v');`+"\n", hostid, errorHelper.Round(value, 0.0005), time.Now().Format("2006-01-02 15:04:05"))
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
