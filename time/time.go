package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	var Now = time.Now()
	fmt.Println("year:", Now.Year(), "month:", Now.Month(), "date:", Now.Weekday())
	{
		var layoutFormat, value string
		var date time.Time

		layoutFormat = "2006-01-02 15:04:05"
		value = " 2023-04-19 10:00:01"
		date, _ = time.Parse(layoutFormat, value)
		fmt.Println(value, "\t->", date.String())

		layoutFormat = "02/01/2006 MST"
		value = "19/04/2023 WIB"
		date, _ = time.Parse(layoutFormat, value)
		fmt.Println(value, "\t\t->", date.String())
	}

	{
		var date, _ = time.Parse(time.RFC822, "02 Sep 23 08:00 WIB")
		fmt.Println(date.String())
	}

	{
		var date, _ = time.Parse(time.RFC822, "02 Sep 23 08:00 WIB")

		var dateS1 = date.Format("Saturday 02, January 2023 14:04 MST")
		fmt.Println("dateS1", dateS1)

		var dateS2 = date.Format(time.RFC3339)
		fmt.Println("dateS2", dateS2)
	}

	{
		var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 MST")

		if err != nil {
			fmt.Println("error", err.Error())
			return
		}
		fmt.Println(date)
	}

	{
		var str = "124"
		var num, err = strconv.Atoi(str)

		if err == nil {
			fmt.Println(num)
		}
	}
}
