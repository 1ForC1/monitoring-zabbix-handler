package errorHelper

import (
	"fmt"
	"strings"
	loghelper "zabbix-data-handler/logHelper"
)

// PrintError Вывод ошибки в консоль и запись в логи
func PrintError(e error) {
	if e != nil {
		fmt.Println("\n************** " + e.Error() + " **************\n")
		loghelper.WriteLogs(strings.ReplaceAll("************** "+e.Error()+" **************\n", "pq: ", ""))
	}
}

// Round Функция округления до тысячных
func Round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}
