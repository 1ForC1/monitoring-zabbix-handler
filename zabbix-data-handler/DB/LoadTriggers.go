package DB

import (
	"fmt"
	"reflect"
	"strconv"
	"zabbix-data-handler/apiRequests"
	"zabbix-data-handler/errorHelper"
	loghelper "zabbix-data-handler/logHelper"
	"zabbix-data-handler/structs"
)

// LoadTriggers Метод получения триггеров из api zabbix
func LoadTriggers() {
	fmt.Printf("Загрузка триггеров(на неисправности)...\n")
	loghelper.WriteLogs("Загрузка триггеров(на неисправности)...")
	for i := 0; i < reflect.ValueOf(structs.Hosts.Result).Len(); i++ {
		if apiRequests.GetTriggers(structs.Hosts.Result[i].Hostid) {

			fmt.Printf("Triggerid: %v, Expression: %v, Description: %v, Priority: %v, HostId: %v\n", structs.Triggers.Result[0].Triggerid, structs.Triggers.Result[0].Expression, structs.Triggers.Result[0].Description, structs.Triggers.Result[0].Priority, structs.Hosts.Result[i].Hostid)
			loghelper.WriteLogs(fmt.Sprintf("Triggerid: %v, Expression: %v, Description: %v, Priority: %v, HostId: %v", structs.Triggers.Result[0].Triggerid, structs.Triggers.Result[0].Expression, structs.Triggers.Result[0].Description, structs.Triggers.Result[0].Priority, structs.Hosts.Result[i].Hostid))

			hostid, err := strconv.ParseInt(structs.Hosts.Result[i].Hostid, 10, 64)
			errorHelper.PrintError(err)
			triggerid, err := strconv.ParseInt(structs.Triggers.Result[0].Triggerid, 10, 64)
			errorHelper.PrintError(err)

			Query += fmt.Sprintf(`insert into "triggers" ("triggerid", "triggers_expression", "description", "priority", "hostid") values (%d , '%v', '%v', '%v', %d);`+"\n", triggerid, structs.Triggers.Result[0].Expression, structs.Triggers.Result[0].Description, structs.Triggers.Result[0].Priority, hostid)
			errorHelper.PrintError(err)
		} else {
			fmt.Printf("Сработавших триггеров по хосту %v %v не обнаружено\n", structs.Hosts.Result[i].Hostid, structs.Hosts.Result[i].HostName)
			loghelper.WriteLogs(fmt.Sprintf("Сработавших триггеров по хосту %v %v не обнаружено", structs.Hosts.Result[i].Hostid, structs.Hosts.Result[i].HostName))
		}
	}

}
