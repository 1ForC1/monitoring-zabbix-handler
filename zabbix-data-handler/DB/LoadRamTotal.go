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

func LoadRamTotal() {
	fmt.Printf("Загрузка данных об общем количестве оперативной памяти...\n")
	loghelper.WriteLogs("Загрузка данных об общем количестве оперативной памяти...")
	body := apiRequests.GetInfoTemplate("GET", "vm.memory.size[total]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.RamTotal)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.RamTotal.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.RamTotal.Result[i].HostId {
					loghelper.PrintHostParameters(structs.RamTotal.Result[i].HostId, structs.RamTotal.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.RamTotal.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "memory_total" ("hostid", "value_memory_total", "time_memory_total") values (%d , %v , '%v');`+"\n", hostid, structs.RamTotal.Result[i].Value, time.Now().Format("2006-01-02 15:04:05"))
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
