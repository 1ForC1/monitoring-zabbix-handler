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

func LoadCpuUtilIowait() {
	fmt.Printf("Загрузка данных заргуженности процессора(ожидание в процентах)...\n")
	loghelper.WriteLogs("Загрузка данных заргуженности процессора(ожидание в процентах)...")
	body := apiRequests.GetInfoTemplate("GET", "system.cpu.util[,iowait]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.CpuUtilIowait)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.CpuUtilIowait.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.CpuUtilIowait.Result[i].HostId {
					loghelper.PrintHostParameters(structs.CpuUtilIowait.Result[i].HostId, structs.CpuUtilIowait.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.CpuUtilIowait.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					value, err := strconv.ParseFloat(structs.CpuUtilIowait.Result[i].Value, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "cpu_Iowait" ("hostid", "value_cpu_Iowait", "time_cpu_Iowait") values (%d , %v , '%v');`+"\n", hostid, errorHelper.Round(value, 0.0005), time.Now().Format("2006-01-02 15:04:05"))
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
