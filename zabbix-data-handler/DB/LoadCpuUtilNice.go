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

func LoadCpuUtilNice() {
	fmt.Printf("Загрузка данных заргуженности процессора(при выполнении на пользовательском уровне с хорошим приоритетом в процентах)...\n")
	loghelper.WriteLogs("Загрузка данных заргуженности процессора(при выполнении на пользовательском уровне с хорошим приоритетом в процентах)...")
	body := apiRequests.GetInfoTemplate("GET", "system.cpu.util[,nice]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.CpuUtilNice)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.CpuUtilNice.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.CpuUtilNice.Result[i].HostId {
					loghelper.PrintHostParameters(structs.CpuUtilNice.Result[i].HostId, structs.CpuUtilNice.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.CpuUtilNice.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					value, err := strconv.ParseFloat(structs.CpuUtilNice.Result[i].Value, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "cpu_nice" ("hostid", "value_cpu_nice", "time_cpu_nice") values (%d , %v , '%v');`+"\n", hostid, errorHelper.Round(value, 0.0005), time.Now().Format("2006-01-02 15:04:05"))
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
