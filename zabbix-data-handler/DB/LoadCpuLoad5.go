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

// LoadCpuLoad5 - метод для загрузки данных заргуженности процессора(среднее за 5 мин)
func LoadCpuLoad5() {
	fmt.Printf("Загрузка данных заргуженности процессора(среднее за 5 мин)...\n")
	loghelper.WriteLogs("Загрузка данных заргуженности процессора(среднее за 5 мин)...")
	body := apiRequests.GetInfoTemplate("GET", "system.cpu.load[percpu,avg5]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.CpuLoad5)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.CpuLoad5.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.CpuLoad5.Result[i].HostId {
					loghelper.PrintHostParameters(structs.CpuLoad5.Result[i].HostId, structs.CpuLoad5.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.CpuLoad5.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					value, err := strconv.ParseFloat(structs.CpuLoad5.Result[i].Value, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "percpu_avg5" ("hostid", "value_percpu_avg5", "time_percpu_avg5") values (%d , %v , '%v');`+"\n", hostid, errorHelper.Round(value, 0.0005), time.Now().Format("2006-01-02 15:04:05"))
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
