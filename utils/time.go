package utils

import (
	"fmt"
	"strconv"
	"time"
)

func TvHourTimeUnix() (int64,int64) {
	nowMinUnix := time.Now().Unix()
	endTime := nowMinUnix - nowMinUnix % 60
	startTime := endTime - 180
	return startTime, endTime
}


func SeveDayUnix() [][]int64 {
	sivetime := time.Now().AddDate(0, 0, -7)
	sivedays := time.Date(sivetime.Year(), sivetime.Month(), sivetime.Day(), 0, 0, 0, 0, time.Local)
	sivedayunix := sivedays.Unix()
	count := sivedayunix + 7 * 86400
	var oneDay [][]int64
	for x := sivedayunix; x < sivedayunix + 86400;x+=60 {
		var y  []int64
		for i := x; i < count; i+=86400 {
			y = append(y,i)
		}
		oneDay = append(oneDay,y)
	}
	return oneDay
}

func Todaydate() string {

	t := time.Now()
	nowday := fmt.Sprintf("%s_%02d_%s", strconv.Itoa(t.Year()), t.Month(), strconv.Itoa(t.Day()))
	return nowday

}