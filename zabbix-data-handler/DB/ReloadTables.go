package DB

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"zabbix-data-handler/apiRequests"
	"zabbix-data-handler/errorHelper"
	loghelper "zabbix-data-handler/logHelper"
	"zabbix-data-handler/structs"
)

// Query Прогоняемый запрос
var Query string

// ReloadTables Удаление данных в таблицах и загрузка данных из zabbix
func ReloadTables() error {
	//Удаление всех данных в таблицах
	Query = `delete from "cpu_interrupt";delete from "cpu_Iowait";delete from "cpu_nice";
delete from "cpu_softirq";delete from "cpu_steal";delete from "cpu_system";delete from "cpu_util_idle";
delete from "cpu_util_user";delete from "memory_available";delete from "memory_total";
delete from "percpu_avg1";delete from "percpu_avg5";delete from "percpu_avg15";delete from "size_free";
delete from "size_total"; delete from "triggers";delete from "host";`
	//delete from "triggers";

	//Загрузка хостов в скрипт
	LoadHosts()
	//Загрузка данных заргуженности процессора(среднее за 15 мин)
	LoadCpuLoad15()
	//Загрузка данных заргуженности процессора(среднее за 5 мин)
	LoadCpuLoad5()
	//Загрузка данных заргуженности процессора(среднее за 1 мин)
	LoadCpuLoad1()
	//Загрузка данных заргуженности процессора(пользователем в процентах)
	LoadCpuUtilUser()
	//Загрузка данных заргуженности процессора(в простое в процентах)
	LoadCpuUtilIdle()
	//Загрузка данных заргуженности процессора(прерывания в процентах)
	LoadCpuUtilInterrupt()
	//Загрузка данных заргуженности процессора(ожидание в процентах)
	LoadCpuUtilIowait()
	//Загрузка данных заргуженности процессора(при выполнении на пользовательском уровне с хорошим приоритетом в процентах)
	LoadCpuUtilNice()
	//Загрузка данных заргуженности процессора(прогаммные прерывания в процентах)
	LoadCpuUtilSoftirq()
	//Загрузка данных заргуженности процессора(время в течении которого виртуальная машина не получает ресурсы процессора в процентах)
	LoadCpuUtilSteal()
	//Загрузка данных заргуженности процессора(системой в процентах)
	LoadCpuUtilSystem()
	//Загрузка данных о свободном месте на диске
	LoadDiskSpaceFree()
	//Загрузка данных об общем количестве места на диске
	LoadDiskSpaceTotal()
	//Загрузка данных о свободной оперативной памяти
	LoadRamFree()
	//Загрузка данных об общем количестве оперативной памяти
	LoadRamTotal()
	//Загрузка данных о триггерах
	LoadTriggers()

	//Выполнение скрипта очищения и загрузки данных в БД
	fmt.Printf("\nВыполнение загрузки данных в БД...")
	loghelper.WriteLogs("Выполнение загрузки данных в БД...")
	fmt.Printf(Query)
	loghelper.WriteLogs(Query)
	_, err = db.Exec(Query)
	errorHelper.PrintError(err)
	return err
}

// LoadHosts Метод получения хостов из api zabbix
func LoadHosts() {
	fmt.Printf("Загрузка хостов...\n")
	loghelper.WriteLogs("Загрузка хостов...")
	apiRequests.GetHost()
	for i := 0; i < reflect.ValueOf(structs.Hosts.Result).Len(); i++ {
		fmt.Printf("Hostid: %v, HostName: %v, HostInterfaces: %s\n", structs.Hosts.Result[i].Hostid, structs.Hosts.Result[i].HostName, structs.Hosts.Result[i].HostIP)
		loghelper.WriteLogs(fmt.Sprintf("Hostid: %v, HostName: %v, HostInterfaces: %s", structs.Hosts.Result[i].Hostid, structs.Hosts.Result[i].HostName, structs.Hosts.Result[i].HostIP))
		hostid, err := strconv.ParseInt(structs.Hosts.Result[i].Hostid, 10, 64)
		errorHelper.PrintError(err)

		//Конвертирование данных ip хостов в формат записи в базу данных
		hostInterfaces := strings.Replace(fmt.Sprintf("%s", structs.Hosts.Result[i].HostIP), "{", "", -1)
		hostInterfaces = strings.Replace(hostInterfaces, "}", "", -1)
		hostInterfaces = strings.Replace(hostInterfaces, "[", "", -1)
		hostInterfaces = strings.Replace(hostInterfaces, "]", "", -1)
		hostInterfaces = strings.Replace(hostInterfaces, " ", ",", -1)
		hostInterfaces = "{" + hostInterfaces + "}"

		//Добавление в строку запроса
		Query += fmt.Sprintf(`insert into "host"("hostid", "name_host", "ip_host") values(%d , '%s', '%s');`+"\n", hostid, structs.Hosts.Result[i].HostName, hostInterfaces)
	}
}
