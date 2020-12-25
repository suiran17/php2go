package php2go

import (
	"fmt"
	"testing"
	"time"
)

// 不好解析
// 似乎不准确, Strtotime, 里面解析应该使用 ParseInLocation, 从而获取当时时区时间对应的unix时间戳
func TestStrtoTime(t *testing.T) {
	// str := "12/05/2020 14:15:05"
	//2016-01-02 23:04:05
	timestamp, err := Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05")
	// timestamp, err := Strtotime("12/05/2020 14:15:05", str)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(timestamp)
}

func GettimeLayout() string {
	return "2006-01-02 15:04:05"
}

// 字符串转为时间戳
func TestStrtostamp(t *testing.T) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	time.Now().In(cstSh)
	
	str := "2020-12-25 14:32:55"
	
	// p, _ := time.Parse("2006-01-02 15:04:05", str)
	p, _ := time.ParseInLocation("2006-01-02 15:04:05", str, cstSh)
	// p, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	fmt.Println(p)        // 2020-12-25 14:32:55 +0000 UTC  , 这里的时间错误, 不应该是 utc 时间, 应该是 东八区的时间
	fmt.Println(p.Unix()) // 不准确, 是utc 时间 1608906775 2020/12/25 22:32:55
	
	fmt.Println(time.Now().Unix()) // 准确
	
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 准确
}

func TestUtcTime(t *testing.T) {
	// now := Date.now()
	// timeLocal := time.Local
	
	fmt.Println(time.Local)
}