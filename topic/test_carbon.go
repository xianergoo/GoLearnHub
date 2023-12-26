package main

import (
	"fmt"

	"github.com/golang-module/carbon/v2"
)

func main() {
	// 今天此刻
	fmt.Sprintf("%s", carbon.Now()) // 2020-08-05 13:14:15
	carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
	// 今天日期
	carbon.Now().ToDateString() // 2020-08-05
	// 今天时间
	carbon.Now().ToTimeString() // 13:14:15
	// 指定时区的今天此刻
	carbon.Now(carbon.NewYork).ToDateTimeString()               // 2020-08-05 14:14:15
	carbon.SetTimezone(carbon.NewYork).Now().ToDateTimeString() // 2020-08-05 14:14:15
	// 今天秒级时间戳
	carbon.Now().Timestamp()           // 1596604455
	carbon.Now().TimestampWithSecond() // 1596604455
	// 今天毫秒级时间戳
	carbon.Now().TimestampWithMillisecond() // 1596604455000
	// 今天微秒级时间戳
	carbon.Now().TimestampWithMicrosecond() // 1596604455000000
	// 今天纳秒级时间戳
	carbon.Now().TimestampWithNanosecond() // 1596604455000000000

	// 昨天此刻
	fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
	carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
	// 昨天日期
	carbon.Yesterday().ToDateString() // 2020-08-04
	// 昨天时间
	carbon.Yesterday().ToTimeString() // 13:14:15
	// 指定日期的昨天此刻
	carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
	// 指定时区的昨天此刻
	carbon.Yesterday(Carbon.NewYork).ToDateTimeString()               // 2020-08-04 14:14:15
	carbon.SetTimezone(Carbon.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 14:14:15
	// 昨天秒级时间戳
	carbon.Yesterday().Timestamp()           // 1596518055
	carbon.Yesterday().TimestampWithSecond() // 1596518055
	// 昨天毫秒级时间戳
	carbon.Yesterday().TimestampWithMillisecond() // 1596518055000
	// 昨天微秒级时间戳
	carbon.Yesterday().TimestampWithMicrosecond() // 1596518055000000
	// 昨天纳秒级时间戳
	carbon.Yesterday().TimestampWithNanosecond() // 1596518055000000000

	// 明天此刻
	fmt.Sprintf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
	carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
	// 明天日期
	carbon.Tomorrow().ToDateString() // 2020-08-06
	// 明天时间
	carbon.Tomorrow().ToTimeString() // 13:14:15
	// 指定日期的明天此刻
	carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
	// 指定时区的明天此刻
	carbon.Tomorrow(Carbon.NewYork).ToDateTimeString()               // 2020-08-06 14:14:15
	carbon.SetTimezone(Carbon.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 14:14:15
	// 明天秒级时间戳
	carbon.Tomorrow().Timestamp()           // 1596690855
	carbon.Tomorrow().TimestampWithSecond() // 1596690855
	// 明天毫秒级时间戳
	carbon.Tomorrow().TimestampWithMillisecond() // 1596690855000
	// 明天微秒级时间戳
	carbon.Tomorrow().TimestampWithMicrosecond() // 1596690855000000
	// 明天纳秒级时间戳
	carbon.Tomorrow().TimestampWithNanosecond() // 1596690855000000000

	// 从秒级时间戳创建 Carbon 实例
	carbon.CreateFromTimestamp(-1).ToDateTimeString()                       // 1970-01-01 07:59:59
	carbon.CreateFromTimestamp(-1, carbon.Tokyo).ToDateTimeString()         // 1970-01-01 08:59:59
	carbon.CreateFromTimestamp(0).ToDateTimeString()                        // 1970-01-01 08:00:00
	carbon.CreateFromTimestamp(0, carbon.Tokyo).ToDateTimeString()          // 1970-01-01 09:00:00
	carbon.CreateFromTimestamp(1596604455).ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// 从毫秒级时间戳创建 Carbon 实例
	carbon.CreateFromTimestamp(1596604455000).ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// 从微秒级时间戳创建 Carbon 实例
	carbon.CreateFromTimestamp(1596604455000000).ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// 从纳级时间戳创建 Carbon 实例
	carbon.CreateFromTimestamp(1596604455000000000).ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455000000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

	// 从年月日时分秒创建 Carbon 实例
	carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// 从年月日创建 Carbon 实例(时分秒默认为当前时分秒)
	carbon.CreateFromDate(2020, 8, 5).ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.CreateFromDate(2020, 8, 5, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)
	carbon.CreateFromTime(13, 14, 15).ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.CreateFromTime(13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

	carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString()                     // 2020-08-05 13:14:15
	carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
	carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString()           // 2020-08-05 13:14:15
	carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString()       // 2020-08-05 14:14:15

}
