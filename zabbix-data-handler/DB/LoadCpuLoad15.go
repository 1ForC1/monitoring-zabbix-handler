package DB

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"zabbix-data-handler/apiRequests"
	"zabbix-data-handler/errorHelper"
	"zabbix-data-handler/logHelper"
	"zabbix-data-handler/structs"
)

func LoadCpuLoad15() {
	fmt.Printf("Загрузка данных заргуженности процессора(среднее за 15 мин)...\n")
	loghelper.WriteLogs("Загрузка данных заргуженности процессора(среднее за 15 мин)...")
	body := apiRequests.GetInfoTemplate("GET", "system.cpu.load[percpu,avg15]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.CpuLoad15)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.CpuLoad15.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.CpuLoad15.Result[i].HostId {
					loghelper.PrintHostParameters(structs.CpuLoad15.Result[i].HostId, structs.CpuLoad15.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.CpuLoad15.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					value, err := strconv.ParseFloat(structs.CpuLoad15.Result[i].Value, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "percpu_avg15" ("hostid", "value_percpu_avg15", "time_percpu_avg15") values (%d , %v , '%v');`+"\n", hostid, errorHelper.Round(value, 0.0005), time.Now().Format("2006-01-02 15:04:05"))
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
