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

func LoadRamFree() {
	fmt.Printf("Загрузка данных о свободной оперативной памяти...\n")
	loghelper.WriteLogs("Загрузка данных о свободной оперативной памяти...")
	body := apiRequests.GetInfoTemplate("GET", "vm.memory.size[available]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.RamFree)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.RamFree.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.RamFree.Result[i].HostId {
					loghelper.PrintHostParameters(structs.RamFree.Result[i].HostId, structs.RamFree.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.RamFree.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "memory_available" ("hostid", "value_memory_available", "time_memory_available") values (%d , %v , '%v');`+"\n", hostid, structs.RamFree.Result[i].Value, time.Now().Format("2006-01-02 15:04:05"))
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
