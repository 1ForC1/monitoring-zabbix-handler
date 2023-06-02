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

func LoadDiskSpaceFree() {
	fmt.Printf("Загрузка данных о свободном месте на диске...\n")
	loghelper.WriteLogs("Загрузка данных о свободном месте на диске...")
	body := apiRequests.GetInfoTemplate("GET", "vfs.fs.size[/,free]", nil)
	if body != nil {
		err = json.Unmarshal(body, &structs.DiskSpaceFree)
		errorHelper.PrintError(err)
		for i := 0; i < reflect.ValueOf(structs.DiskSpaceFree.Result).Len(); i++ {
			for j := 0; j < reflect.ValueOf(structs.Hosts.Result).Len(); j++ {
				if structs.Hosts.Result[j].Hostid == structs.DiskSpaceFree.Result[i].HostId {
					loghelper.PrintHostParameters(structs.DiskSpaceFree.Result[i].HostId, structs.DiskSpaceFree.Result[i].Value)
					hostid, err := strconv.ParseInt(structs.DiskSpaceFree.Result[i].HostId, 10, 32)
					errorHelper.PrintError(err)
					Query += fmt.Sprintf(`insert into "size_free" ("hostid", "value_size_free", "time_size_free") values (%d , %v , '%v');`+"\n", hostid, structs.DiskSpaceFree.Result[i].Value, time.Now().Format("2006-01-02 15:04:05"))
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
