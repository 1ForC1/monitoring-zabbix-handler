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

func LoadDiskSpaceTotal() {
	fmt.Printf("Загрузка данных об общем количестве места на диске...\n")
	loghelper.WriteLogs("Загрузка данных об общем количестве места на диске...")
	body := apiRequests.GetInfoTemplate("GET", "vfs.fs.size[/,total]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.DiskSpaceTotal)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.DiskSpaceTotal.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.DiskSpaceTotal.Result[i].HostId {
					loghelper.PrintHostParameters(structs.DiskSpaceTotal.Result[i].HostId, structs.DiskSpaceTotal.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.DiskSpaceTotal.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "size_total" ("hostid", "value_size_total", "time_size_total") values (%d , %v , '%v');`+"\n", hostid, structs.DiskSpaceTotal.Result[i].Value, time.Now().Format("2006-01-02 15:04:05"))
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
