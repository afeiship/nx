package nx

import (
	"fmt"
	"github.com/afeiship/nx/lib"
	"testing"
	"time"
)

func TestStrftime(t *testing.T) {
	nx.Hi()
	// create a time : 2019-01-01 00:00:00
	currentTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	// fmt.Println("当前时间  : ", currentTime)
	// fmt.Println("当前时间字符串: ", currentTime.String())
	// fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	// fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))

	// start test date
	format := "YYYY-MM-DD"
	res := nx.Strftime(format, currentTime)
	if res != "2019-01-01" {
		t.Error("Strftime failed!")
	}

	// start test time
	format = "HH:mm:SS"
	res = nx.Strftime(format, currentTime)
	if res != "00:00:00" {
		t.Error("Strftime failed!")
	}

	// start test datetime
	format = "YYYY-MM-DD HH:mm:SS"
	res = nx.Strftime(format, currentTime)
	if res != "2019-01-01 00:00:00" {
		t.Error("Strftime failed!")
	}

	// test dateHook
	format = "date"
	res = nx.Strftime(format, currentTime)
	if res != "2019-01-01" {
		t.Error("Strftime failed!")
	}

	// test timeHook
	// start test time
	format = "time"
	res = nx.Strftime(format, currentTime)
	if res != "00:00:00" {
		t.Error("Strftime failed!")
	}

	// test datetimeHook
	// start test datetime
	format = "datetime"
	res = nx.Strftime(format, currentTime)
	if res != "2019-01-01 00:00:00" {
		t.Error("Strftime failed!")
	}

	// test dbHook
	format = "db"
	res = nx.Strftime(format, currentTime)
	fmt.Println(res)
	if res != "20190101_000000" {
		t.Error("Strftime failed!")
	}
}
